package entity

type Temperature struct {
	Celsius    float64 `json:"celsius"`
	Fahrenheit float64 `json:"fahrenheit"`
	Kelvin     float64 `json:"kelvin"`
}

func NewTemperature(celsius float64, fahrenheit float64) Temperature {
	return Temperature{
		Celsius:    celsius,
		Fahrenheit: fahrenheit,
		Kelvin:     celsius + 273.15,
	}
}
