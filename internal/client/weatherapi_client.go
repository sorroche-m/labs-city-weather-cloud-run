package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/sorroche-m/weather/internal/model"
)

type WeatherAPIClient struct {
	baseURL string
	apiKey  string
	client  *http.Client
}

func NewWeatherAPIClient() *WeatherAPIClient {
	apiKey := os.Getenv("WEATHER_API_KEY")

	return &WeatherAPIClient{
		baseURL: "https://api.weatherapi.com/v1",
		apiKey:  apiKey,
		client:  &http.Client{},
	}
}

func (c *WeatherAPIClient) GetWeather(city string) (*model.WeatherAPIResponse, error) {
	params := url.Values{}
	params.Add("key", c.apiKey)
	params.Add("q", city)
	params.Add("aqi", "no")

	url := fmt.Sprintf("%s/current.json?%s", c.baseURL, params.Encode())

	resp, err := c.client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error fetching weather: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var weatherResp model.WeatherAPIResponse
	if err := json.NewDecoder(resp.Body).Decode(&weatherResp); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	return &weatherResp, nil
}
