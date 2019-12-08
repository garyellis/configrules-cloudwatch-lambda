package aws

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
)

// PutMetricData writes the config metric to cloudwatch
func PutMetricData(svc *cloudwatch.CloudWatch, namespace, metricname, unit string, value int) {
	input := &cloudwatch.PutMetricDataInput{
		Namespace: aws.String(namespace),
		MetricData: []*cloudwatch.MetricDatum{
			&cloudwatch.MetricDatum{
				MetricName: aws.String(metricname),
				Unit:       aws.String(unit),
				Value:      aws.Float64(float64(value)),
				Dimensions: []*cloudwatch.Dimension{},
			},
		},
	}
	_, err := svc.PutMetricData(input)
	if err != nil {
		log.Printf("[aws/cloudwatch] %s", err)
	}
}

// PutMetricAlarm creates the config rule cloudwatch metric alarm
func PutMetricAlarm(svc *cloudwatch.CloudWatch, namespace string, metricname string, name string, alarmActions []string) error {
	input := &cloudwatch.PutMetricAlarmInput{
		AlarmName:          aws.String(name),
		MetricName:         aws.String(metricname),
		Namespace:          aws.String(namespace),
		Period:             aws.Int64(300),
		ComparisonOperator: aws.String(cloudwatch.ComparisonOperatorLessThanThreshold),
		Threshold:          aws.Float64(1.0),
		Statistic:          aws.String(cloudwatch.StatisticMaximum),
		EvaluationPeriods:  aws.Int64(1),
		ActionsEnabled:     aws.Bool(true),
		AlarmActions:       aws.StringSlice(alarmActions),
	}
	if _, err := svc.PutMetricAlarm(input); err != nil {
		log.Printf("[aws/cloudwatch] %s", err)
	}

	return nil
}
