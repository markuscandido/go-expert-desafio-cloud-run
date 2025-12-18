package weatherapi

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/markuscandido/go-expert-desafio-cloud-run/internal/domain"
)

const baseURL = "https://api.weatherapi.com/v1"

// weatherAPIResponse represents the API response
type weatherAPIResponse struct {
	Location struct {
		Name    string `json:"name"`
		Region  string `json:"region"`
		Country string `json:"country"`
	} `json:"location"`
	Current struct {
		TempC float64 `json:"temp_c"`
		TempF float64 `json:"temp_f"`
	} `json:"current"`
}

// Client implements WeatherRepository using WeatherAPI
type Client struct {
	httpClient *http.Client
	apiKey     string
}

// NewClient creates a new WeatherAPI client
func NewClient(apiKey string) *Client {
	return &Client{
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
		apiKey: apiKey,
	}
}

// GetTemperature returns the current temperature in Celsius for a given city
func (c *Client) GetTemperature(ctx context.Context, city string) (float64, error) {
	endpoint := fmt.Sprintf("%s/current.json?key=%s&q=%s&aqi=no",
		baseURL,
		c.apiKey,
		url.QueryEscape(city),
	)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return 0, fmt.Errorf("%w: %v", domain.ErrExternalService, err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return 0, fmt.Errorf("%w: %v", domain.ErrExternalService, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusBadRequest || resp.StatusCode == http.StatusNotFound {
		return 0, domain.ErrWeatherNotFound
	}

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("%w: status code %d", domain.ErrExternalService, resp.StatusCode)
	}

	var result weatherAPIResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return 0, fmt.Errorf("%w: %v", domain.ErrExternalService, err)
	}

	return result.Current.TempC, nil
}
