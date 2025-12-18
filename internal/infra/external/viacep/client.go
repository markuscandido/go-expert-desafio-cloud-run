package viacep

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/markuscandido/go-expert-desafio-cloud-run/internal/domain"
	"github.com/markuscandido/go-expert-desafio-cloud-run/internal/domain/entity"
)

const baseURL = "https://viacep.com.br/ws"

// viaCEPResponse represents the API response
type viaCEPResponse struct {
	CEP         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	UF          string `json:"uf"`
	Erro        string `json:"erro"`
}

// Client implements LocationRepository using ViaCEP API
type Client struct {
	httpClient *http.Client
}

// NewClient creates a new ViaCEP client
func NewClient() *Client {
	return &Client{
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

// GetLocation returns the location for a given zipcode
func (c *Client) GetLocation(ctx context.Context, zipcode string) (*entity.Location, error) {
	url := fmt.Sprintf("%s/%s/json/", baseURL, zipcode)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", domain.ErrExternalService, err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", domain.ErrExternalService, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%w: status code %d", domain.ErrExternalService, resp.StatusCode)
	}

	var result viaCEPResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("%w: %v", domain.ErrExternalService, err)
	}

	// ViaCEP returns {"erro": "true"} when CEP is not found
	if result.Erro == "true" || result.Localidade == "" {
		return nil, domain.ErrZipcodeNotFound
	}

	return entity.NewLocation(result.Localidade, result.UF), nil
}
