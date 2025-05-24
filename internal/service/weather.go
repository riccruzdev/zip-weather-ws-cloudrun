package service

import (
	"time"

	"github.com/riccruzdev/zip-weather-ws-cloudrun/internal/usecase"
	"github.com/riccruzdev/zip-weather-ws-cloudrun/pkg/cep"
	"github.com/riccruzdev/zip-weather-ws-cloudrun/pkg/weather"
)

type WeatherService struct {
	WeatherUsecase *usecase.WeatherUsecase
}

func NewWeatherService(weatherAPIKey string) *WeatherService {

	cepClient := cep.NewCEPClientImpl(5 * time.Second)
	weatherClient := weather.NewWeatherClientImpl(5 * time.Second)
	weatherUsecase := usecase.NewWeatherUsecase(cepClient, weatherClient)

	return &WeatherService{
		WeatherUsecase: &weatherUsecase,
	}
}
