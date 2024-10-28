package http

import (
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/superdev/ecommerce/gateway/internal/adapters/http/handlers"
	"github.com/superdev/ecommerce/gateway/internal/adapters/http/middlewares"
)

func NewHTTPRouter(
	handlers []handlers.Handler,
	_middlewares []middlewares.Middleware,
) (*gin.Engine, error) {

	engine := gin.Default()

	secured := engine.Group("/")

	// Add middlewares
	for _, middleware := range _middlewares {
		secured.Use(middleware.Handle())
	}
	for _, h := range handlers {
		if h.Pattern() == "/login" {
			engine.POST(h.Pattern(), h.Handle)
			continue
		}
		methods := h.Methods()
		if len(methods) == 0 {
			return nil, errors.New("found a HTTP handler with empty method(verb) list. Please specify at least one method")
		} else if len(methods) == 1 {
			method := methods[0]
			switch strings.ToUpper(method) {
			case "GET":
				secured.GET(
					h.Pattern(),
					middlewares.Authenticate(h),
					h.Handle,
				)
			case "PATCH":
				secured.PATCH(
					h.Pattern(),
					middlewares.Authenticate(h),
					h.Handle,
				)
			case "PUT":
				secured.PUT(
					h.Pattern(),
					middlewares.Authenticate(h),
					h.Handle,
				)

			case "POST":
				secured.POST(
					h.Pattern(),
					middlewares.Authenticate(h),
					h.Handle,
				)
			case "DELETE":
				secured.DELETE(
					h.Pattern(),
					middlewares.Authenticate(h),
					h.Handle,
				)
			}

		} else {
			secured.Match(
				h.Methods(),
				h.Pattern(),
				middlewares.Authenticate(h),
				h.Handle,
			)
		}
	}

	return engine, nil
}
