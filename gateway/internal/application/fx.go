package application

import (
	"github.com/superdev/ecommerce/gateway/internal/application/command/user"
	"github.com/superdev/ecommerce/gateway/internal/application/service"
	"github.com/superdev/ecommerce/gateway/internal/command"
	"go.uber.org/fx"
)

var Module = fx.Module("app.service",
	fx.Provide(
		service.NewProductService,
		service.NewOrderService,
		service.NewUserService,
		//
		service.NewPasswordService,
		service.ProvidePasswordConfig,
		//
		command.AsCobraCommand(user.NCreateUserCommand),
		command.AsCobraCommand(user.NewRootUserCommand),
	),
)
