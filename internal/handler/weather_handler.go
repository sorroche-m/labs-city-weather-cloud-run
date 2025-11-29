package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/sorroche-m/weather/internal/model"
	"github.com/sorroche-m/weather/internal/service"
)

type WeatherHandler struct {
	service *service.WeatherService
}

func NewWeatherHandler() *WeatherHandler {
	return &WeatherHandler{
		service: service.NewWeatherService(),
	}
}

func (h *WeatherHandler) GetWeather(w http.ResponseWriter, r *http.Request) {
	cep := chi.URLParam(r, "cep")

	weather, err := h.service.GetWeatherByCEP(cep)
	if err != nil {
		h.handleError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(weather)
}

func (h *WeatherHandler) handleError(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")

	errMsg := err.Error()
	var statusCode int
	var message string

	message = errMsg

	switch errMsg {
	case "invalid zipcode":
		statusCode = http.StatusUnprocessableEntity
	case "can not find zipcode":
		statusCode = http.StatusNotFound
	default:
		statusCode = http.StatusInternalServerError
		message = "internal server error"
	}

	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(model.ErrorResponse{Message: message})
}
