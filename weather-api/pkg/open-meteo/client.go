package countriesnow

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

const (
	getCityCoordinates = "https://geocoding-api.open-meteo.com/v1/search?name=Kyiv&count=1&language=en&format=json"
)

type Client struct{}

func NewClient() *Client {
	return &Client{}
}

type CoordinatesResponse struct {
	Error       bool          `json:"error"`
	Message     string        `json:"msg"`
	Coordinates Coordinates `json:"data"`
}

type CoordinatesRequest struct {
	Country string `json:"country"`
	City    string `json:"city"`
}

type Coordinates struct {
	Lat float64 `json:"latitude"`
	Lon float64 `json:"longitude"`
}

func (c *Client) Coordinates(_ context.Context) (*Coordinates, error) {
	resp, err := http.Get(getCityCoordinates)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var dto
	if err = json.Unmarshal(body, dto); err != nil {
		return nil, err
	}

	if dto.Error {
		return nil, errors.New(dto.Message)
	}

	return &dto.Coordinates, nil
}
