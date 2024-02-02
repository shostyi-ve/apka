package weather_prediction_api

import (
	"go.uber.org/fx"
)

var Module = fx.Module("weather",
	fx.Provide(New),
)
