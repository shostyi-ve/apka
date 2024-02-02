package config

import (
	"github.com/joeshaw/envdecode"
	"github.com/pkg/errors"
)

type Config struct {
	HTTPServerPort        string `env:"HTTP_SERVER_PORT"`
	WeatherPredictionHost string `env:"WEATHER_PREDICTION_HOST"`
}

func NewConfig() (*Config, error) {
	config := &Config{}

	if err := envdecode.Decode(config); err != nil {
		return nil, errors.WithStack(err)
	}

	return config, nil
}
