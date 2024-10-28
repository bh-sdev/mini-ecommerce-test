package config

import "go.uber.org/fx"

var Module = fx.Module("app.config",
	fx.Provide(
		fx.Annotate(
			NewConfig,
		),
	),
)
