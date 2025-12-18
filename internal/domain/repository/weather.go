package repository

import (
	"context"
)

// WeatherRepository defines the interface for weather lookup
type WeatherRepository interface {
	// GetTemperature returns the current temperature in Celsius for a given city
	// Returns domain.ErrWeatherNotFound if weather data is not available
	GetTemperature(ctx context.Context, city string) (float64, error)
}
