package env

import (
	"os"
	"strings"
)

func GetEnv(key string) string {
	return os.Getenv(key)
}

func GetEnvOrDefault(key string, def string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return def
}

func GetBoolEnvOrDefault(key string, def bool) bool {
	if value, ok := os.LookupEnv(key); ok {
		value = strings.TrimSpace(value)
		value = strings.ToLower(value)
		if value == "true" || value == "1" || value == "yes" || value == "y" || value == "on" {
			return true
		}
		if value == "false" || value == "0" || value == "no" || value == "n" || value == "off" {
			return false
		}
		return value != ""
	}
	return def
}
