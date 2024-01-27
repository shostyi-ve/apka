package server

import (
	"github.com/gofiber/fiber"
	"github.com/shostyi-ve/apka/weather/config"
	"go.uber.org/fx"
)

var Module = fx.Module("server",
	fx.Provide(fiber.New),
	fx.Invoke(func(lifecycle fx.Lifecycle, app *fiber.App, cfg *config.Config) {
		lifecycle.Append(
			fx.StartHook(func() {
				go app.Listen(cfg.HTTPServerPort)
			}),
		)
	}),
)
