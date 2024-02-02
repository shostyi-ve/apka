package weather_prediction_api

import (
	"bytes"
	"context"
	"fmt"
	"github.com/shostyi-ve/apka/weather/config"
	"io/ioutil"
	"net/http"
)

type (
	Client struct {
		baseURL string
	}
)

const pathPrediction = "/api/weather/get_prediction_24h"

func New(config *config.Config) *Client {
	return &Client{config.WeatherPredictionHost}
}

func (c *Client) GetPrediction(ctx context.Context, history string) (string, error) {
	url := c.baseURL + pathPrediction

	requestBody := []byte(history)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, bytes.NewBuffer(requestBody))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		return string(result), nil
	}

	return "", fmt.Errorf("HTTP request failed with status code: %d", resp.StatusCode)
}
