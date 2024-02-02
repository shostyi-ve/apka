package weather

import (
	"go.uber.org/fx"
)

var Module = fx.Provide(NewWeatherService)
