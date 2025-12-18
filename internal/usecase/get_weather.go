package usecase

import (
	"context"
	"regexp"

	"github.com/markuscandido/go-expert-desafio-cloud-run/internal/domain"
	"github.com/markuscandido/go-expert-desafio-cloud-run/internal/domain/entity"
	"github.com/markuscandido/go-expert-desafio-cloud-run/internal/domain/repository"
)

// zipcodeRegex validates Brazilian CEP format (8 digits)
var zipcodeRegex = regexp.MustCompile(`^\d{8}$`)

// GetWeatherByZipcodeUseCase handles the business logic for getting weather by zipcode
type GetWeatherByZipcodeUseCase struct {
	locationRepo repository.LocationRepository
	weatherRepo  repository.WeatherRepository
}

// NewGetWeatherByZipcodeUseCase creates a new use case instance
func NewGetWeatherByZipcodeUseCase(
	locationRepo repository.LocationRepository,
	weatherRepo repository.WeatherRepository,
) *GetWeatherByZipcodeUseCase {
	return &GetWeatherByZipcodeUseCase{
		locationRepo: locationRepo,
		weatherRepo:  weatherRepo,
	}
}

// Execute runs the use case
// 1. Validates the zipcode format
// 2. Gets location from zipcode
// 3. Gets temperature from location
// 4. Returns weather with all temperature units
func (uc *GetWeatherByZipcodeUseCase) Execute(ctx context.Context, zipcode string) (*entity.Weather, error) {
	// Validate zipcode format (8 digits)
	if !zipcodeRegex.MatchString(zipcode) {
		return nil, domain.ErrInvalidZipcode
	}

	// Get location from zipcode
	location, err := uc.locationRepo.GetLocation(ctx, zipcode)
	if err != nil {
		return nil, err
	}

	// Get temperature from location
	tempC, err := uc.weatherRepo.GetTemperature(ctx, location.City)
	if err != nil {
		return nil, err
	}

	// Create weather with all temperature units
	return entity.NewWeather(tempC), nil
}
