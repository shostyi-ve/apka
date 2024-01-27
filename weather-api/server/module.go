package server

import (
	"fmt"

	"github.com/gofiber/fiber"
	"github.com/shostyi-ve/apka/weather/config"
	"go.uber.org/fx"
)

var Module = fx.Module("server",
	fx.Provide(fiber.New),
	fx.Invoke(func(app *fiber.App, cfg *config.Config) {
		app.Listen(fmt.Sprintf(":%s", cfg.HTTPServerPort))
	}),
)
