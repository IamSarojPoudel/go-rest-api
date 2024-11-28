package config

import "os"

type Config struct {
	ServerAddress string
	DatabaseURL   string
}

func LoadConfig() *Config {
	return &Config{
		ServerAddress: getEnv("SERVER_ADDRESS", ":8080"),
		DatabaseURL:   getEnv("DATABASE_URL", "host=localhost user=youruser password=yourpassword dbname=yourdb port=5432 sslmode=disable"),
	}
}

func getEnv(key, fallback string) string {
	if value, exist := os.LookupEnv(key); exist {
		return value
	}
	return fallback
}
