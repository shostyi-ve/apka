package db

import (
	"context"
	"go.uber.org/fx"
)

var Module = fx.Provide(
	context.TODO,
	NewDBConf,
	GetDBConnect,
)
