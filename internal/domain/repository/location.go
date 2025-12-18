package repository

import (
	"context"

	"github.com/markuscandido/go-expert-desafio-cloud-run/internal/domain/entity"
)

// LocationRepository defines the interface for location lookup
type LocationRepository interface {
	// GetLocation returns the location for a given zipcode
	// Returns domain.ErrZipcodeNotFound if zipcode doesn't exist
	GetLocation(ctx context.Context, zipcode string) (*entity.Location, error)
}
