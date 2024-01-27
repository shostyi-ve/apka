package container

import (
	"github.com/shostyi-ve/apka/weather/config"
	app "github.com/shostyi-ve/apka/weather/internal"
	"github.com/shostyi-ve/apka/weather/server"

	"go.uber.org/fx"
)

func Start() {
	fx.New(
		server.Module,
		app.Module,
		config.Module,
	)
}