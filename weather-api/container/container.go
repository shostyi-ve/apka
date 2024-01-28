package container

import (
	"github.com/shostyi-ve/apka/weather/config"
	"github.com/shostyi-ve/apka/weather/db"
	app "github.com/shostyi-ve/apka/weather/internal"
	"github.com/shostyi-ve/apka/weather/internal/repositories"
	"github.com/shostyi-ve/apka/weather/server"
	"go.uber.org/fx"
)

func Build() *fx.App {
	return fx.New(
		config.Module,

		db.Module,
		repositories.Module,

		server.Module,

		app.Module,
	)
}
