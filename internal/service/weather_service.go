package service

import (
	"fmt"
	"regexp"

	"github.com/sorroche-m/weather/internal/client"
	"github.com/sorroche-m/weather/internal/model"
)

type WeatherService struct {
	viaCEPClient     *client.ViaCEPClient
	weatherAPIClient *client.WeatherAPIClient
}

func NewWeatherService() *WeatherService {
	return &WeatherService{
		viaCEPClient:     client.NewViaCEPClient(),
		weatherAPIClient: client.NewWeatherAPIClient(),
	}
}

func (s *WeatherService) GetWeatherByCEP(cep string) (*model.WeatherResponse, error) {
	if !s.isValidCEP(cep) {
		return nil, fmt.Errorf("invalid zipcode")
	}

	location, err := s.viaCEPClient.GetLocation(cep)
	if err != nil {
		return nil, fmt.Errorf("can not find zipcode")
	}

	weather, err := s.weatherAPIClient.GetWeather(location.Localidade)
	if err != nil {
		return nil, fmt.Errorf("error fetching weather: %w", err)
	}

	tempC := weather.Current.TempC
	tempF := s.celsiusToFahrenheit(tempC)
	tempK := s.celsiusToKelvin(tempC)

	return &model.WeatherResponse{
		TempC: tempC,
		TempF: tempF,
		TempK: tempK,
	}, nil
}

func (s *WeatherService) isValidCEP(cep string) bool {
	matched, _ := regexp.MatchString(`^\d{8}$`, cep)
	return matched
}

func (s *WeatherService) celsiusToFahrenheit(celsius float64) float64 {
	return celsius*1.8 + 32
}

func (s *WeatherService) celsiusToKelvin(celsius float64) float64 {
	return celsius + 273
}
