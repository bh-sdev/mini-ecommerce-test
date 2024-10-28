package logger

import "go.uber.org/fx"

var Module = fx.Module("app.logger",
	fx.Provide(
		fx.Annotate(
			NewZapLogger,
		),
	),
)
