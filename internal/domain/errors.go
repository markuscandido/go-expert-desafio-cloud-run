package domain

import "errors"

var (
	// ErrInvalidZipcode indicates the zipcode format is invalid
	ErrInvalidZipcode = errors.New("invalid zipcode")

	// ErrZipcodeNotFound indicates the zipcode was not found
	ErrZipcodeNotFound = errors.New("can not find zipcode")

	// ErrWeatherNotFound indicates weather data was not found for the location
	ErrWeatherNotFound = errors.New("weather not found")

	// ErrExternalService indicates an error with external service communication
	ErrExternalService = errors.New("external service error")
)
