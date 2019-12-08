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
  name_prefix = "config-metrics-lambda"
  policy = data.aws_iam_policy_document.permission_policy.json
}

resource "aws_iam_role" "iam_role" {
  name_prefix        = "config-metrics-lambda"
  description        = "aws config custom metrics"
  assume_role_policy = data.aws_iam_policy_document.trust_policy.json
}

resource "aws_iam_role_policy_attachment" "iam_role" {
  role       = aws_iam_role.iam_role.name
  policy_arn = aws_iam_policy.iam_policy.arn
}

#### lambda function
data "external" "release" {
  program = [""]
  query   = {
    url = ""
    out = "../function.zip"
  }
}
