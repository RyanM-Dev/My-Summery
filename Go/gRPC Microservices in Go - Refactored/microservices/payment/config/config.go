package config

import (
	"log"
	"os"
	"strconv"
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

// GetTLSCertDir returns TLS_CERT_DIR, or "" when TLS is off.
// The folder must hold ca-cert.pem, server-cert.pem, and server-key.pem.
func GetTLSCertDir() string {
	return os.Getenv("TLS_CERT_DIR")
}

// getEnvironmentValue stops the service at startup when a required value is missing.
func getEnvironmentValue(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("%s environment variable is missing.", key)
	}
	return value
}
