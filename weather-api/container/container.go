package container

import (
	"github.com/shostyi-ve/apka/weather/config"
	app "github.com/shostyi-ve/apka/weather/internal"
	weather_prediction_api "github.com/shostyi-ve/apka/weather/internal/infrustructure/external/weather-prediction-api"
	"github.com/shostyi-ve/apka/weather/internal/repositories"
	"github.com/shostyi-ve/apka/weather/internal/service/weather"
	"github.com/shostyi-ve/apka/weather/pkg/countriesnow"
	"github.com/shostyi-ve/apka/weather/pkg/db"
	"github.com/shostyi-ve/apka/weather/pkg/open_meteo"
	"github.com/shostyi-ve/apka/weather/pkg/server"
	"go.uber.org/fx"
)

func Build() *fx.App {
	return fx.New(
		app.Module,
		config.Module,
		repositories.Module,
		server.Module,
		db.Module,
		countriesnow.Module,
		weather.Module,
		weather_prediction_api.Module,
		open_meteo.Module,
	)
}
