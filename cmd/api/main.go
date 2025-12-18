package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/markuscandido/go-expert-desafio-cloud-run/configs"
	"github.com/markuscandido/go-expert-desafio-cloud-run/internal/infra/external/viacep"
	"github.com/markuscandido/go-expert-desafio-cloud-run/internal/infra/external/weatherapi"
	"github.com/markuscandido/go-expert-desafio-cloud-run/internal/infra/web"
	"github.com/markuscandido/go-expert-desafio-cloud-run/internal/usecase"
)

func main() {
	// Setup structured logging
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	// Load configuration
	cfg := configs.LoadConfig()

	if cfg.WeatherAPIKey == "" {
		slog.Warn("WEATHER_API_KEY not set, weather lookups will fail")
	}

	// Initialize dependencies
	viaCEPClient := viacep.NewClient()
	weatherAPIClient := weatherapi.NewClient(cfg.WeatherAPIKey)

	// Initialize use case
	getWeatherUseCase := usecase.NewGetWeatherByZipcodeUseCase(viaCEPClient, weatherAPIClient)

	// Initialize handler
	weatherHandler := web.NewWeatherHandler(getWeatherUseCase)

	// Setup routes
	mux := http.NewServeMux()
	mux.HandleFunc("/weather/", weatherHandler.GetWeather)
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	// Start server
	addr := fmt.Sprintf(":%s", cfg.ServerPort)
	slog.Info("starting server", "addr", addr)

	if err := http.ListenAndServe(addr, mux); err != nil {
		slog.Error("server failed", "error", err)
		os.Exit(1)
	}
}
