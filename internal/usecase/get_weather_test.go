package usecase

import (
	"context"
	"errors"
	"testing"

	"github.com/markuscandido/go-expert-desafio-cloud-run/internal/domain"
	"github.com/markuscandido/go-expert-desafio-cloud-run/internal/domain/entity"
)

// MockLocationRepository is a mock for LocationRepository
type MockLocationRepository struct {
	GetLocationFunc func(ctx context.Context, zipcode string) (*entity.Location, error)
}

func (m *MockLocationRepository) GetLocation(ctx context.Context, zipcode string) (*entity.Location, error) {
	return m.GetLocationFunc(ctx, zipcode)
}

// MockWeatherRepository is a mock for WeatherRepository
type MockWeatherRepository struct {
	GetTemperatureFunc func(ctx context.Context, city string) (float64, error)
}

func (m *MockWeatherRepository) GetTemperature(ctx context.Context, city string) (float64, error) {
	return m.GetTemperatureFunc(ctx, city)
}

func TestGetWeatherByZipcodeUseCase_Execute(t *testing.T) {
	tests := []struct {
		name            string
		zipcode         string
		mockLocation    *entity.Location
		mockLocationErr error
		mockTemp        float64
		mockTempErr     error
		wantErr         error
		wantTempC       float64
	}{
		{
			name:         "valid zipcode returns weather",
			zipcode:      "01310100",
			mockLocation: &entity.Location{City: "São Paulo", State: "SP"},
			mockTemp:     28.5,
			wantTempC:    28.5,
		},
		{
			name:    "invalid zipcode format - too short",
			zipcode: "123",
			wantErr: domain.ErrInvalidZipcode,
		},
		{
			name:    "invalid zipcode format - letters",
			zipcode: "1234567a",
			wantErr: domain.ErrInvalidZipcode,
		},
		{
			name:    "invalid zipcode format - too long",
			zipcode: "123456789",
			wantErr: domain.ErrInvalidZipcode,
		},
		{
			name:            "zipcode not found",
			zipcode:         "99999999",
			mockLocationErr: domain.ErrZipcodeNotFound,
			wantErr:         domain.ErrZipcodeNotFound,
		},
		{
			name:         "weather not found",
			zipcode:      "01310100",
			mockLocation: &entity.Location{City: "São Paulo", State: "SP"},
			mockTempErr:  domain.ErrWeatherNotFound,
			wantErr:      domain.ErrWeatherNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			locationRepo := &MockLocationRepository{
				GetLocationFunc: func(ctx context.Context, zipcode string) (*entity.Location, error) {
					if tt.mockLocationErr != nil {
						return nil, tt.mockLocationErr
					}
					return tt.mockLocation, nil
				},
			}

			weatherRepo := &MockWeatherRepository{
				GetTemperatureFunc: func(ctx context.Context, city string) (float64, error) {
					if tt.mockTempErr != nil {
						return 0, tt.mockTempErr
					}
					return tt.mockTemp, nil
				},
			}

			uc := NewGetWeatherByZipcodeUseCase(locationRepo, weatherRepo)
			weather, err := uc.Execute(context.Background(), tt.zipcode)

			if tt.wantErr != nil {
				if !errors.Is(err, tt.wantErr) {
					t.Errorf("Execute() error = %v, wantErr %v", err, tt.wantErr)
				}
				return
			}

			if err != nil {
				t.Errorf("Execute() unexpected error = %v", err)
				return
			}

			if weather.TempC != tt.wantTempC {
				t.Errorf("Execute() TempC = %v, want %v", weather.TempC, tt.wantTempC)
			}
		})
	}
}
