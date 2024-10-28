package database

import "go.uber.org/fx"

var Module = fx.Module("app.database",
	fx.Provide(
		fx.Annotate(
			NewDB,
		),
	),
)
