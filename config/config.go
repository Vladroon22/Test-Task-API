package config

import "os"

type Config struct {
	Addr_PORT string
	DB        string
}

func CreateConfig() *Config {
	return &Config{
		Addr_PORT: getEnv("addr_port", ":5000"),
		DB:        getEnv("DB", ""),
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
