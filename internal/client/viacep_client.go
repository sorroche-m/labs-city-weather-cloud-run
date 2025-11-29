package client

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sorroche-m/weather/internal/model"
)

type ViaCEPClient struct {
	baseURL string
	client  *http.Client
}

func NewViaCEPClient() *ViaCEPClient {
	return &ViaCEPClient{
		baseURL: "https://viacep.com.br/ws",
		client:  &http.Client{},
	}
}

func (c *ViaCEPClient) GetLocation(cep string) (*model.ViaCEPResponse, error) {
	url := fmt.Sprintf("%s/%s/json/", c.baseURL, cep)

	resp, err := c.client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error fetching CEP: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var viaCEPResp model.ViaCEPResponse
	if err := json.NewDecoder(resp.Body).Decode(&viaCEPResp); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	if viaCEPResp.Erro {
		return nil, fmt.Errorf("CEP not found")
	}

	return &viaCEPResp, nil
}
