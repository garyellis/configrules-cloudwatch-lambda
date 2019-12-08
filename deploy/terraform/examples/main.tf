variable "env_rule_whitelist" {}
variable "env_create_alarms" {}
variable "env_alarm_arns" {}

module "configrules_cloudwatch_lambda" {
  source = "../"

  env_rule_whitelist = var.env_rule_whitelist
  env_create_alarms  = var.env_create_alarms
  env_alarm_arns     = var.env_alarm_arns 
}
