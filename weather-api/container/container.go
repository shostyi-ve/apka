package container

import (
	"github.com/shostyi-ve/apka/weather/config"
	app "github.com/shostyi-ve/apka/weather/internal"
	"github.com/shostyi-ve/apka/weather/internal/repositories"
	"github.com/shostyi-ve/apka/weather/pkg/countriesnow"
	"github.com/shostyi-ve/apka/weather/pkg/db"
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
	)
}
