package cmd

import (
	"fmt"
	"log"
	"strconv"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
	"github.com/garyellis/config-metrics-lambda/pkg/aws"
	"github.com/garyellis/config-metrics-lambda/pkg/config"
)

var namespace = "config-metrics-lambda"
var configRuleStatuses []aws.ConfigRuleStatus

func ConfigRuleStatusesToCloudWatch(cfg *config.Config) error {
	ruleStatuses, _ := aws.CollectConfigRuleStatuses(cfg.RulesWhiteList)
	for _, rule := range ruleStatuses {
		log.Printf("[cmd] config rule compliance status for %s is compliancestatus: %s\n", rule.Name, rule.ComplianceStatus)
	}
	PutConfigRuleMetrics(ruleStatuses, cfg.CreateAlarms, cfg.AlarmActionArns)
	return nil
}

func PutConfigRuleMetrics(ruleStatuses []aws.ConfigRuleStatus, createAlarms bool, alarmActions []string) {
	log.Printf("[cmd] alarm creation flag is set to %s", strconv.FormatBool(createAlarms))
	svc := cloudwatch.New(session.New())
	for _, rule := range ruleStatuses {
		// write the metrics
		log.Printf("[cmd] writing %s metrics to cloudwatch", rule.Name)
		aws.PutMetricData(svc, namespace, fmt.Sprintf("compliance-%s", rule.Name), "None", rule.IsCompliantMetric())

		// create the metric alarm
		if createAlarms {
			log.Printf("[cmd] creating cloudwatch alarm for %s", rule.Name)
			if err := aws.PutMetricAlarm(svc, namespace, fmt.Sprintf("compliance-%s", rule.Name), fmt.Sprintf("compliance-%s", rule.Name), alarmActions); err != nil {
				log.Printf("[cmd] %s", err)
			}
		}
	}
}
