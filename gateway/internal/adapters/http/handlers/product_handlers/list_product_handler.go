package producthandlers

import (
	"fmt"
	_http "net/http"

	"github.com/gin-gonic/gin"
	"github.com/superdev/ecommerce/gateway/internal/adapters/http/handlers"
	"github.com/superdev/ecommerce/gateway/internal/config"
	"github.com/superdev/ecommerce/gateway/internal/domain/service"
)

type listProductHandler struct {
	handlers.AbstractAdminHandler
	productService service.ProductService
}

func NewListProductHandler(
	productService service.ProductService,
	conf *config.Config,
) handlers.Handler {
	return &listProductHandler{
		productService: productService,
	}
}

func (s *listProductHandler) Pattern() string {
	return "/products"
}

func (s *listProductHandler) Methods() []string {
	return []string{"GET"}
}

func (s *listProductHandler) Handle(c *gin.Context) {

	products, err := s.productService.ListProducts()
	if err != nil {
		c.JSON(_http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": fmt.Sprintf("Could not list Products: %s", err.Error()),
		})
		return
	}

	c.JSON(_http.StatusOK, gin.H{
		"status":  "success",
		"message": ("Product listd successfully"),
		"data":    products, //product,
	})

}
