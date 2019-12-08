# configrules-cloudwatch-lambda
This lambda provides cloudwatch alarms integration for aws config rules. Its intendended use is passive monitoring and alerting on config rule compliance statuses.


## Environment variables

| Name | Description | Default | Required |
|------|-------------|:-------:|:--------:|
|  `RULE_WHITELIST` | A list of comma separated config rules | `""` | no |
|  `CREATE_ARNS` | enable or disable cloudwatch alarm creation | `"true"` | no |
|  `ALARM_ARNS` | A list of comma separated notification ARNs | `""` | no |
