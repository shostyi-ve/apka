package app

import (
	"github.com/gofiber/fiber"
	"github.com/shostyi-ve/apka/weather/internal/handler"
	"go.uber.org/fx"
)

var Module = fx.Module("weather",
	fx.Provide(
		handler.NewWeatherHandler,
		handler.NewGeoHandler,
	),
	fx.Invoke(func(h *handler.WeatherHandler, app *fiber.App) {
		h.RegisterRoutes(app)
	}),
	fx.Invoke(func(h *handler.GeoHandler, app *fiber.App) {
		h.RegisterRoutes(app)
	}),
)
