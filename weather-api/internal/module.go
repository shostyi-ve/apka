package app

import (
	"github.com/gofiber/fiber"
	"github.com/shostyi-ve/apka/weather/internal/handler"
	"go.uber.org/fx"
)

const addr = "localhost:8080"

var Module = fx.Module("example",
	fx.Provide(
		handler.NewWeatherHandler,
	),
	fx.Invoke(func(h *handler.WeatherHandler, app *fiber.App) {
		h.RegisterRoutes(app)
	}),
)
