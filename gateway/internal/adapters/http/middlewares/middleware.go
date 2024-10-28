package middlewares

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type Middleware interface {
	Handle() gin.HandlerFunc
}

func AsHttpMiddleware(f any, anns ...fx.Annotation) any {
	anns = append(anns,
		fx.As(new(Middleware)),
		fx.ResultTags(`group:"app.http.middleware"`))

	return fx.Annotate(
		f,
		anns...,
	)
}
