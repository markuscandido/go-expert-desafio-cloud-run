package entity

import "math"

// Weather represents temperature in multiple units
type Weather struct {
	TempC float64 `json:"temp_C"`
	TempF float64 `json:"temp_F"`
	TempK float64 `json:"temp_K"`
}

// NewWeather creates a Weather from Celsius temperature
// Converts to Fahrenheit: F = C * 1.8 + 32
// Converts to Kelvin: K = C + 273
func NewWeather(tempC float64) *Weather {
	return &Weather{
		TempC: round(tempC),
		TempF: round(celsiusToFahrenheit(tempC)),
		TempK: round(celsiusToKelvin(tempC)),
	}
}

// celsiusToFahrenheit converts Celsius to Fahrenheit
// Formula: F = C * 1.8 + 32
func celsiusToFahrenheit(celsius float64) float64 {
	return celsius*1.8 + 32
}

// celsiusToKelvin converts Celsius to Kelvin
// Formula: K = C + 273
func celsiusToKelvin(celsius float64) float64 {
	return celsius + 273
}

// round rounds to 1 decimal place
func round(value float64) float64 {
	return math.Round(value*10) / 10
}
