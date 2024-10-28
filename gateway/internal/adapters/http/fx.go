package http

import (
	"github.com/superdev/ecommerce/gateway/internal/adapters/http/handlers"
	authhandlers "github.com/superdev/ecommerce/gateway/internal/adapters/http/handlers/auth_handlers"
	orderhandlers "github.com/superdev/ecommerce/gateway/internal/adapters/http/handlers/order_handlers"
	producthandlers "github.com/superdev/ecommerce/gateway/internal/adapters/http/handlers/product_handlers"
	"github.com/superdev/ecommerce/gateway/internal/adapters/http/middlewares"
	"go.uber.org/fx"
)

var Module = fx.Module("http_adapter",
	// fx.Invoke(func(*cobra.Command) {}),
	fx.Provide(
		NewHTTPRunCmd,
		NewHTTPServer,
		fx.Annotate(
			NewHTTPRouter,
			fx.ParamTags(`group:"app.http.handler"`, `group:"app.http.middleware"`),
		),

		/* ---- HANDLERS  ------     */
		// AUTH
		handlers.AsHttpHandler(
			authhandlers.NewRegisterAuthHandler,
		),
		handlers.AsHttpHandler(
			authhandlers.NewLoginAuthHandler,
		),
		// PRODUCTS
		handlers.AsHttpHandler(
			producthandlers.NewCreateProductHandler,
		),
		handlers.AsHttpHandler(
			producthandlers.NewGetProductHandler,
		),
		handlers.AsHttpHandler(
			producthandlers.NewListProductHandler,
		),
		handlers.AsHttpHandler(
			producthandlers.NewUpdateProductHandler,
		),
		handlers.AsHttpHandler(
			producthandlers.NewDeleteProductHandler,
		),
		// ORDERS
		handlers.AsHttpHandler(
			orderhandlers.NewCreateOrderHandler,
		),
		handlers.AsHttpHandler(
			orderhandlers.NewGetOrderHandler,
		),
		handlers.AsHttpHandler(
			orderhandlers.NewListOrderHandler,
		),
		handlers.AsHttpHandler(
			orderhandlers.NewDeleteOrderHandler,
		),
		/* ---- MIDDLEWARES ------*/
		middlewares.AsHttpMiddleware(
			middlewares.NewJWTMiddleware,
		),
	),
)
