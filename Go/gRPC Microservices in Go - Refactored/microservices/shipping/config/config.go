package config

import (
	"log"
	"os"
	"strconv"
	"strings"
)

// GetEnv returns ENV, such as "development" or "production".
func GetEnv() string {
	return getEnvironmentValue("ENV")
}

func GetDataSourceURL() string {
	return getEnvironmentValue("DATA_SOURCE_URL")
}

func GetApplicationPort() int {
	portStr := getEnvironmentValue("APPLICATION_PORT")
	port, err := strconv.Atoi(portStr)
	if err != nil {
		log.Fatalf("port: %s is invalid", portStr)
	}
	return port
}

// GetKafkaBrokers returns KAFKA_BROKERS split on commas, such as
// ["kafka:19092"].
func GetKafkaBrokers() []string {
	return strings.Split(getEnvironmentValue("KAFKA_BROKERS"), ",")
}

// getEnvironmentValue stops the service at startup when a required value is missing.
func getEnvironmentValue(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("%s environment variable is missing.", key)
	}
	return value
}
