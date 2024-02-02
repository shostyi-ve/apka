package weather

import (
	"context"
	weather_prediction_api "github.com/shostyi-ve/apka/weather/internal/infrustructure/external/weather-prediction-api"
	"github.com/shostyi-ve/apka/weather/pkg/open_meteo"
	"time"
)

type WeatherService struct {
	openMeteoClient  *open_meteo.Client
	predictionClient *weather_prediction_api.Client
}

func NewWeatherService(openMeteoClient *open_meteo.Client, predictionClient *weather_prediction_api.Client) *WeatherService {
	return &WeatherService{openMeteoClient: openMeteoClient, predictionClient: predictionClient}
}

func (w *WeatherService) GetPrediction(ctx context.Context, city string) (string, error) {
	coords, err := w.openMeteoClient.Coordinates(ctx, open_meteo.CoordinatesRequest{City: city})
	if err != nil {
		return "", err
	}

	history, err := w.openMeteoClient.History(ctx, open_meteo.HistoryRequest{
		Coordinates: open_meteo.Coordinates{
			Lat: coords.Lat,
			Lon: coords.Lon,
		},
		StartDate: time.Now().In(time.FixedZone("GMT", 0)).Add(-7 * 24 * time.Hour),
		EndData:   time.Now().In(time.FixedZone("GMT", 0)),
	})
	if err != nil {
		return "", err
	}

	return w.predictionClient.GetPrediction(ctx, history.Data)
}

func (w *WeatherService) GetHistory(ctx context.Context, start time.Time, end time.Time, city string) (string, error) {
	coords, err := w.openMeteoClient.Coordinates(ctx, open_meteo.CoordinatesRequest{City: city})
	if err != nil {
		return "", err
	}

	history, err := w.openMeteoClient.History(ctx, open_meteo.HistoryRequest{
		Coordinates: open_meteo.Coordinates{
			Lat: coords.Lat,
			Lon: coords.Lon,
		},
		StartDate: start,
		EndData:   end,
	})
	if err != nil {
		return "", err
	}

	return history.Data, nil
}
