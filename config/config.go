package config

import (
	"os"
)

//DatabaseConfig properties
type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Name     string
	Password string
}

//Config project config
type Config struct {
	ServiceName    string
	Port           string
	Database       DatabaseConfig
	LightstepToken string
}

//New returns a new instance of config
func New() *Config {
	return &Config{
		ServiceName:    getEnv("SERVICE_NAME", "transaction_service"),
		Port:           getEnv("PORT", "8080"),
		LightstepToken: getEnv("LIGHTSTEP_ACCESS_TOKEN", ""),
		Database: DatabaseConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "3000"),
			User:     getEnv("DB_USER", "root"),
			Password: getEnv("DB_PASSWORD", "password"),
			Name:     getEnv("DB_NAME", "database"),
		},
	}
}

func getEnv(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
