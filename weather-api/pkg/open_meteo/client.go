package open_meteo

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

const (
	geocodingAPIURL = "https://geocoding-api.open-meteo.com/v1/search"
	historyAPIURL   = "https://archive-api.open-meteo.com/v1/archive?latitude=%f&longitude=%f&start_date=2024-01-16&end_date=2024-01-30&hourly=temperature_2m,relative_humidity_2m,cloud_cover,wind_speed_10m,wind_direction_10m,snowfall,rain,weather_code,visibility,is_day"
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
	City string `json:"city"`
}

type HistoryResponse struct {
	Data        string
	DateCreated time.Time
}

type HistoryRequest struct {
	Coordinates Coordinates
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

func (c *Client) History(_ context.Context, req HistoryRequest) (*HistoryResponse, error) {
	apiURL := fmt.Sprintf(historyAPIURL, req.Coordinates.Lat, req.Coordinates.Lon)

	resp, err := http.Get(apiURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var dto HistoryResponse
	dto.Data = string(body)
	dto.DateCreated = time.Now()

	if err = json.Unmarshal(body, &dto); err != nil {
		return nil, err
	}

	return &dto, nil
}
