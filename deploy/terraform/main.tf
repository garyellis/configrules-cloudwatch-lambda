locals {
  name = var.name_prefix == "" ? "configrules-cloudwatch-lambda" : format("%s-configrules-cloudwatch-lambda", var.name_prefix)
}

data "aws_iam_policy_document" "trust_policy" {
  statement {
    actions = ["sts:AssumeRole"]
    effect  = "Allow"
    principals {
      type        = "Service"
      identifiers =  ["lambda.amazonaws.com"]
    }
  }
}

data "aws_iam_policy_document" "permission_policy" {
  statement {
    sid     = "EventsWriteCWL"
    effect  = "Allow"
    actions = [
      "logs:CreateLogGroup",
      "logs:CreateLogStream",
      "logs:PutLogEvents",
      "logs:PutSubscriptionFilter",
      "config:Get*",
      "config:Describe*",
      "config:Deliver*",
      "config:List*",
      "tag:GetResources",
      "tag:GetTagKeys",
      "cloudtrail:DescribeTrails",
      "cloudtrail:GetTrailStatus",
      "cloudtrail:LookupEvents",
      "cloudwatch:*"
    ]
    resources = ["*"]
  }
}

resource "aws_iam_policy" "iam_policy" {
  name_prefix = local.name
  policy = data.aws_iam_policy_document.permission_policy.json
}

resource "aws_iam_role" "iam_role" {
  name_prefix        = local.name
  description        = "aws config custom metrics"
  assume_role_policy = data.aws_iam_policy_document.trust_policy.json
}

resource "aws_iam_role_policy_attachment" "iam_role" {
  role       = aws_iam_role.iam_role.name
  policy_arn = aws_iam_policy.iam_policy.arn
}

#### lambda function
data "external" "get_function_zip" {
  program = ["${path.module}/get_function_zip.sh"]
  query   = {
    url = "https://github.com/garyellis/configrules-cloudwatch-lambda/releases/download/v0.1.0/function.zip"
    out = "./function.zip"
  }
}

resource "aws_lambda_function" "lambda" {
  filename         = data.external.get_function_zip.result.path
  source_code_hash = filebase64sha256(data.external.get_function_zip.result.path)
  function_name    = local.name
  role             = aws_iam_role.iam_role.arn
  handler          = "main"
  runtime          = "go1.x"
  tags             = merge(map("Name", "configrules-cloudwatch-lambda"), var.tags)
  environment {
    variables = {
      ALARM_ARNS     = var.env_alarm_arns
      CREATE_ALARMS  = var.env_create_alarms
      RULE_WHITELIST = var.env_rule_whitelist
    }
  }
}

### event triggers
resource "aws_cloudwatch_event_rule" "lambda_cron" {
  name_prefix         = local.name
  description         = ""
  is_enabled          = true
  schedule_expression = "rate(5 minutes)"
  tags                = var.tags
}

resource "aws_cloudwatch_event_target" "lambda" {
  rule      = aws_cloudwatch_event_rule.lambda_cron.name
  target_id = "configrules-cloudwatch-lambda"
  arn       = aws_lambda_function.lambda.arn
}

resource "aws_lambda_permission" "lambda" {
  statement_id  = local.name
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.lambda.function_name
  principal     = "events.amazonaws.com"
  source_arn    = aws_cloudwatch_event_rule.lambda_cron.arn
}
