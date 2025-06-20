package config

import "os"

type AppConfig struct {
	Database string
	Host     string
	Port     string

	FixedAuthToken string
}

func LoadConfig() AppConfig {
	return AppConfig{
		Database:       os.Getenv("MONGO_DATABASE"),
		Host:           os.Getenv("MONGO_HOST"),
		Port:           os.Getenv("MONGO_PORT"),
		FixedAuthToken: os.Getenv("FIXED_AUTH_TOKEN"),
	}
}
