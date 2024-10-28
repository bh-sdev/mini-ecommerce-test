package data

import (
	gormrepo "github.com/superdev/ecommerce/gateway/internal/data/repository/gorm_repo"
	"github.com/superdev/ecommerce/gateway/internal/data/repository/grpc"
	"go.uber.org/fx"
)

var Module = fx.Module("app.data",
	fx.Provide(
		fx.Annotate(
			gormrepo.NewGormUserRepository,
		),
		fx.Annotate(
			grpc.NewGrpcProductRepository,
		),
		fx.Annotate(
			grpc.NewGrpcOrderRepository,
		),
	),
)
