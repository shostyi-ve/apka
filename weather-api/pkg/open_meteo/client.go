package open_meteo

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const (
	geocodingAPIURL = "https://geocoding-api.open-meteo.com/v1/search"
)

type Client struct{}

func NewClient() *Client {
	return &Client{}
}

type CoordinatesResponse struct {
	Results          []Coordinates `json:"results"`
	GenerationTimeMS float64       `json:"generationtime_ms"`
}

type CoordinatesRequest struct {
	Country string `json:"country"`
	City    string `json:"city"`
}

type Coordinates struct {
	Lat float64 `json:"latitude"`
	Lon float64 `json:"longitude"`
}

func (c *Client) Coordinates(_ context.Context, req CoordinatesRequest) (*Coordinates, error) {
	apiURL := fmt.Sprintf("%s?name=%s&count=1&language=en&format=json", geocodingAPIURL, url.QueryEscape(req.City))

	resp, err := http.Get(apiURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var dto CoordinatesResponse
	if err = json.Unmarshal(body, &dto); err != nil {
		return nil, err
	}

	return &dto.Results[0], nil
}
