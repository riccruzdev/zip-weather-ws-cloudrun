package handler

import (
	"encoding/json"
	"net/http"

	"github.com/riccruzdev/zip-weather-ws-cloudrun/internal/service"
)

type WeatherHandler struct {
	WeatherService *service.WeatherService
}

func NewWeatherHandler(weatherService *service.WeatherService) WeatherHandler {
	return WeatherHandler{
		WeatherService: weatherService,
	}
}

func (h *WeatherHandler) GetTemperature(apiKey string, w http.ResponseWriter, r *http.Request) {
	zipcode := r.URL.Query().Get("cep")
	if zipcode == "" {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	weather, err := h.WeatherService.WeatherUsecase.Execute(apiKey, zipcode)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	jsonResponse, err := json.Marshal(weather)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}
