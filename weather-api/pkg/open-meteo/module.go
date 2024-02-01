package countriesnow

import "go.uber.org/fx"

var Module = fx.Module("countriesnow",
	fx.Provide(NewClient),
)
