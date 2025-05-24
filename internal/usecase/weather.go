package usecase

import (
	"github.com/riccruzdev/zip-weather-ws-cloudrun/internal/entity"
	"github.com/riccruzdev/zip-weather-ws-cloudrun/pkg/cep"
	"github.com/riccruzdev/zip-weather-ws-cloudrun/pkg/weather"
)

type WeatherUsecase struct {
	CEPClient     cep.CEPClient
	WeatherClient weather.WeatherClient
}

func NewWeatherUsecase(cepClient cep.CEPClient, weatherClient weather.WeatherClient) WeatherUsecase {
	return WeatherUsecase{
		CEPClient:     cepClient,
		WeatherClient: weatherClient,
	}
}

func (uc *WeatherUsecase) Execute(apiKey string, zipcode string) (entity.Temperature, error) {
	cepResponse, err := uc.CEPClient.GetCityByZipcode(zipcode)
	if err != nil {
		return entity.Temperature{}, err
	}

	weatherResponse, err := uc.WeatherClient.GetTemperatureByCity(apiKey, cepResponse.Localidade)
	if err != nil {
		return entity.Temperature{}, err
	}

	return entity.NewTemperature(weatherResponse.TempC, weatherResponse.TempF), nil
}
