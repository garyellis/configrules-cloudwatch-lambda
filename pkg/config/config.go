package config

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func GetEnv(key, defaultval string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultval
}

func GetEnvBool(key string, defaultval bool) bool {
	if val, err := strconv.ParseBool(GetEnv(key, "")); err == nil {
		return val
	}
	return defaultval
}

func GetEnvSlice(key string, defaultval []string, sep string) []string {
	val := strings.Split(GetEnv(key, ""), sep)
	if len(val) > 0 {
		return val
	}
	return defaultval
}

// Config is the lambda configuration options
type Config struct {
	RulesWhiteList  []string
	CreateAlarms    bool
	AlarmActionArns []string
}

func New() *Config {
	var defaultwhitelist = []string{}
	var defaultalarmarns = []string{}
	var defaultcreatealarms = false

	log.Print("[config] initializing config")

	whitelist := GetEnvSlice("RULE_WHITELIST", defaultwhitelist, ",")
	alarmarns := GetEnvSlice("ALARM_ARNS", defaultalarmarns, ",")
	createalarms := GetEnvBool("CREATE_ALARMS", defaultcreatealarms)

	log.Printf("[config] read environment variable  RULE_WHITELIST: %s", whitelist)
	log.Printf("[config] read environemnt variable ALARM_ARNS: %s", alarmarns)
	log.Printf("[config] read environment variable CREATE_ALARMS: %s", strconv.FormatBool(createalarms))

	return &Config{
		RulesWhiteList:  whitelist,
		CreateAlarms:    createalarms,
		AlarmActionArns: alarmarns,
	}
}
