package grpc

import "go.uber.org/fx"

var Module = fx.Module("grpc_adapter",
	fx.Provide(
		// AsCommand(NewHelloCmd),
		// AsCommand(NewHiCmd),
		fx.Annotate(
			NewProductGrpcClient,
		),
		fx.Annotate(
			NewOrderGrpcClient,
		),
	),
)
