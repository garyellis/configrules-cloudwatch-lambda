package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/garyellis/config-metrics-lambda/pkg/cmd"
	"github.com/garyellis/config-metrics-lambda/pkg/config"
)

var invokeCount = 0

func LambdaHandler() (int, error) {
	invokeCount = invokeCount + 1
	c := config.New()
	cmd.ConfigRuleStatusesToCloudWatch(c)
	return invokeCount, nil
}

func main() {
	mode := config.GetEnv("MODE", "LAMBDA")
	switch mode {
	case "LAMBDA":
		lambda.Start(LambdaHandler)
	case "CLI":
		c := config.New()
		cmd.ConfigRuleStatusesToCloudWatch(c)
	}
}
