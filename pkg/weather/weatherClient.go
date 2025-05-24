package weather

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

const (
	ErrWeatherRequestFail    = "falha ao requisitar weather API"
	ErrWeatherAPIStatus      = "weather API retornou o status code %d"
	ErrWeatherResponseDecode = "falha ao decodificar resposta da API"
	ErrWeatherInvalid        = "localização inválida"
)

type Temperature struct {
	LastUpdated string  `json:"last_updated"`
	TempC       float64 `json:"temp_c"`
	TempF       float64 `json:"temp_f"`
	IsDay       int     `json:"is_day"`
	Condition   struct {
		Text string `json:"text"`
		Icon string `json:"icon"`
		Code int    `json:"code"`
	} `json:"condition"`
	WindMph    float64 `json:"wind_mph"`
	WindKph    float64 `json:"wind_kph"`
	WindDegree int     `json:"wind_degree"`
	WindDir    string  `json:"wind_dir"`
	PressureMb float64 `json:"pressure_mb"`
	PressureIn float64 `json:"pressure_in"`
	PrecipMm   float64 `json:"precip_mm"`
	PrecipIn   float64 `json:"precip_in"`
	Humidity   int     `json:"humidity"`
	Cloud      int     `json:"cloud"`
	FeelslikeC float64 `json:"feelslike_c"`
	FeelslikeF float64 `json:"feelslike_f"`
	VisKm      float64 `json:"vis_km"`
	VisMiles   float64 `json:"vis_miles"`
	Uv         float64 `json:"uv"`
	GustMph    float64 `json:"gust_mph"`
	GustKph    float64 `json:"gust_kph"`
}

type WeatherClient interface {
	GetTemperatureByCity(apiKey string, city string) (Temperature, error)
}

type WeatherClientImpl struct {
	HTTPClient *http.Client
}

func (w *WeatherClientImpl) GetTemperatureByCity(apiKey string, city string) (Temperature, error) {
	url := fmt.Sprintf("https://api.weatherapi.com/v1/current.json?key=%s&q=%s", apiKey, city)
	resp, err := w.HTTPClient.Get(url)

	if err != nil {
		return Temperature{}, errors.New(ErrWeatherRequestFail)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return Temperature{}, fmt.Errorf(ErrWeatherAPIStatus, resp.StatusCode)
	}

	var response Temperature
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return Temperature{}, errors.New(ErrWeatherResponseDecode)
	}

	if response.LastUpdated == "" {
		return Temperature{}, errors.New(ErrWeatherInvalid)
	}

	return response, nil
}

func NewWeatherClientImpl(timeout time.Duration) WeatherClient {
	return &WeatherClientImpl{
		HTTPClient: &http.Client{
			Timeout: timeout,
		},
	}
}
