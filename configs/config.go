package configs

import (
	"os"
)

// Config holds application configuration
type Config struct {
	ServerPort    string
	WeatherAPIKey string
}

// LoadConfig loads configuration from environment variables
func LoadConfig() *Config {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	return &Config{
		ServerPort:    port,
		WeatherAPIKey: os.Getenv("WEATHER_API_KEY"),
	}
}
