package cep

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

const (
	ErrCEPInvalid        = "CEP inválido"
	ErrCEPAPIStatus      = "CEP API retornou o status code %d"
	ErrCEPRequestFail    = "falha ao requisitar CEP API"
	ErrCEPResponseDecode = "falha ao decodificar resposta da API"
)

type CEPResponse struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Unidade     string `json:"unidade"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Estado      string `json:"estado"`
	Regiao      string `json:"regiao"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

type CEPClient interface {
	GetCityByZipcode(zipcode string) (CEPResponse, error)
}

type CEPClientImpl struct {
	HTTPClient *http.Client
}

func (c *CEPClientImpl) GetCityByZipcode(zipcode string) (CEPResponse, error) {
	url := fmt.Sprintf("https://viacep.com.br/ws/%s/json/", zipcode)
	resp, err := c.HTTPClient.Get(url)

	if err != nil {
		return CEPResponse{}, errors.New(ErrCEPRequestFail)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return CEPResponse{}, fmt.Errorf(ErrCEPAPIStatus, resp.StatusCode)
	}

	var response CEPResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return CEPResponse{}, errors.New(ErrCEPResponseDecode)
	}

	if response.Cep == "" {
		return CEPResponse{}, errors.New(ErrCEPInvalid)
	}

	return response, nil
}

func NewCEPClientImpl(timeout time.Duration) CEPClient {
	return &CEPClientImpl{
		HTTPClient: &http.Client{
			Timeout: timeout,
		},
	}
}
