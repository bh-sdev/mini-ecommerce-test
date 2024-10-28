package handlers

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

// import "net/http"

type Handler interface {
	Handle(c *gin.Context)
	Pattern() string
	Methods() []string
	Roles() *[]string
}

type AbstractHandler struct {
}

func (h *AbstractHandler) Roles() *[]string {
	return &[]string{"admin", "user"}
}

type AbstractAdminHandler struct {
}

func (h *AbstractAdminHandler) Roles() *[]string {
	return &[]string{"admin"}
}

type AbstractUserHandler struct {
}

func (h *AbstractUserHandler) Roles() *[]string {
	return &[]string{"user"}
}

func AsHttpHandler(f any, anns ...fx.Annotation) any {
	anns = append(anns,
		fx.As(new(Handler)),
		fx.ResultTags(`group:"app.http.handler"`))

	return fx.Annotate(
		f,
		anns...,
	)
}
