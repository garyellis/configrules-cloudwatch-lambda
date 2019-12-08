package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/configservice"
)

type ConfigRuleStatus struct {
	Name             string
	ComplianceStatus string
	State            string
}

func (r *ConfigRuleStatus) IsCompliantMetric() int {
	var v int
	switch r.ComplianceStatus {
	case "NON_COMPLIANT":
		v = 0
	case "INSUFFICIENT_DATA", "COMPLIANT":
		v = 1
	}
	return v
}

func getRules(svc *configservice.ConfigService, filter []string) ([]*configservice.ConfigRule, error) {
	input := &configservice.DescribeConfigRulesInput{
		ConfigRuleNames: aws.StringSlice(filter),
	}
	output, _ := svc.DescribeConfigRules(input)
	return output.ConfigRules, nil
}

func getRuleComplianceType(rule string, svc *configservice.ConfigService) string {
	var complianceType string
	input := &configservice.DescribeComplianceByConfigRuleInput{
		ConfigRuleNames: []*string{aws.String(rule)},
	}
	output, _ := svc.DescribeComplianceByConfigRule(input)
	for _, rule := range output.ComplianceByConfigRules {
		complianceType = aws.StringValue(rule.Compliance.ComplianceType)
	}
	return complianceType
}

func CollectConfigRuleStatuses(filter []string) ([]ConfigRuleStatus, error) {
	var ruleStatuses []ConfigRuleStatus
	svc := configservice.New(session.New())
	rules, _ := getRules(svc, filter)
	for _, rule := range rules {
		ruleStatus := ConfigRuleStatus{
			Name:             aws.StringValue(rule.ConfigRuleName),
			ComplianceStatus: getRuleComplianceType(aws.StringValue(rule.ConfigRuleName), svc),
			State:            aws.StringValue(rule.ConfigRuleState),
		}
		ruleStatuses = append(ruleStatuses, ruleStatus)
	}
	return ruleStatuses, nil
}
