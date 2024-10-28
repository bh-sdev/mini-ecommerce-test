package producthandlers

import (
	"fmt"
	_http "net/http"

	"github.com/gin-gonic/gin"
	"github.com/superdev/ecommerce/gateway/internal/adapters/http/handlers"
	"github.com/superdev/ecommerce/gateway/internal/config"
	"github.com/superdev/ecommerce/gateway/internal/domain/service"
)

type deleteProductHandler struct {
	handlers.AbstractAdminHandler
	productService service.ProductService
}

func NewDeleteProductHandler(
	productService service.ProductService,
	conf *config.Config,
) handlers.Handler {
	return &deleteProductHandler{
		productService: productService,
	}
}

func (s *deleteProductHandler) Pattern() string {
	return "/products/:id"
}

func (s *deleteProductHandler) Methods() []string {
	return []string{"DELETE"}
}

func (s *deleteProductHandler) Handle(c *gin.Context) {
	id := c.Param("id")
	product, err := s.productService.GetProduct(id)
	if err != nil {
		c.JSON(_http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": fmt.Sprintf("Error fetching Product: %s", err.Error()),
		})
		return
	}
	err = s.productService.Delete(product)
	if err != nil {
		c.JSON(_http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": fmt.Sprintf("Could not delete Product: %s", err.Error()),
		})
		return
	}

	c.JSON(_http.StatusOK, gin.H{
		"status":  "success",
		"message": ("Product deleted successfully"),
		"data":    nil, //product,
	})

}
