package web

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"strings"

	"github.com/markuscandido/go-expert-desafio-cloud-run/internal/domain"
	"github.com/markuscandido/go-expert-desafio-cloud-run/internal/usecase"
)

// ErrorResponse represents an error response
type ErrorResponse struct {
	Message string `json:"message"`
}

// WeatherHandler handles HTTP requests for weather
type WeatherHandler struct {
	useCase *usecase.GetWeatherByZipcodeUseCase
}

// NewWeatherHandler creates a new weather handler
func NewWeatherHandler(useCase *usecase.GetWeatherByZipcodeUseCase) *WeatherHandler {
	return &WeatherHandler{
		useCase: useCase,
	}
}

// GetWeather handles GET /weather/{cep}
func (h *WeatherHandler) GetWeather(w http.ResponseWriter, r *http.Request) {
	// Extract zipcode from URL path
	// Expected format: /weather/{cep}
	path := strings.TrimPrefix(r.URL.Path, "/weather/")
	zipcode := strings.TrimSpace(path)

	if zipcode == "" {
		h.respondError(w, http.StatusBadRequest, "zipcode is required")
		return
	}

	slog.Info("processing weather request", "zipcode", zipcode)

	weather, err := h.useCase.Execute(r.Context(), zipcode)
	if err != nil {
		h.handleError(w, err)
		return
	}

	h.respondJSON(w, http.StatusOK, weather)
}

// handleError maps domain errors to HTTP responses
func (h *WeatherHandler) handleError(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, domain.ErrInvalidZipcode):
		h.respondError(w, http.StatusUnprocessableEntity, "invalid zipcode")
	case errors.Is(err, domain.ErrZipcodeNotFound):
		h.respondError(w, http.StatusNotFound, "can not find zipcode")
	case errors.Is(err, domain.ErrWeatherNotFound):
		h.respondError(w, http.StatusNotFound, "can not find zipcode")
	default:
		slog.Error("internal error", "error", err)
		h.respondError(w, http.StatusInternalServerError, "internal server error")
	}
}

// respondJSON writes a JSON response
func (h *WeatherHandler) respondJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		slog.Error("failed to encode response", "error", err)
	}
}

// respondError writes an error response
func (h *WeatherHandler) respondError(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(ErrorResponse{Message: message})
}
