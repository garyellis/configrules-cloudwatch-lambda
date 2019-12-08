variable "name_prefix" {
  description = "a prefix applied to resources names"
  type        = string
  default     = ""
}

variable "env_rule_whitelist" {
  description = "a comma separated list of configrules to target"
  type        = string
  default     = ""
}

variable "env_create_alarms" {
  description = "create the lambda alarms"
  type        = string
  default     = "true"
}

variable "env_alarm_arns" {
  description = "a comma separated list of alarm notification arns"
  type        = string
  default     = ""
}

variable "tags" {
  description = "A map of tags applied to all taggable resources"
  type        = map(string)
  default     = {}
}
