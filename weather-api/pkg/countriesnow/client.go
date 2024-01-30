package countriesnow

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

const (
	getCountriesURL       = "https://countriesnow.space/api/v0.1/countries/iso"
	getCitiesByCountryURL = "https://countriesnow.space/api/v0.1/countries/cities"
)

type Client struct{}

func NewClient() *Client {
	return &Client{}
}

type AllCountriesResponse struct {
	Error     bool      `json:"error"`
	Message   string    `json:"msg"`
	Countries []Country `json:"data"`
}

type Country struct {
	Name string `json:"name"`
	ISO2 string `json:"iso2"`
	ISO3 string `json:"iso3"`
}

func (c *Client) AllCountries(_ context.Context) ([]Country, error) {
	resp, err := http.Get(getCountriesURL)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var dto = AllCountriesResponse{}
	if err := json.Unmarshal(body, &dto); err != nil {
		return nil, err
	}

	if dto.Error {
		return nil, errors.New(dto.Message)
	}

	return dto.Countries, nil
}

type CitiesByCountryRequest struct {
	CountryName string `json:"country"`
}

type CitiesByCountryResponse struct {
	Error   bool     `json:"error"`
	Message string   `json:"msg"`
	Cities  []string `json:"data"`
}

func (c *Client) CitiesByCountry(_ context.Context, req CitiesByCountryRequest) ([]string, error) {
	reqBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	fmt.Println(string(reqBody))
	resp, err := http.Post(getCitiesByCountryURL, "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var dto = CitiesByCountryResponse{}

	if err := json.Unmarshal(body, &dto); err != nil {
		return nil, err
	}

	if dto.Error {
		return nil, errors.New(dto.Message)
	}

	return dto.Cities, nil
}
