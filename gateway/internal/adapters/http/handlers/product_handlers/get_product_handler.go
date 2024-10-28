package producthandlers

import (
	"fmt"
	_http "net/http"

	"github.com/gin-gonic/gin"
	"github.com/superdev/ecommerce/gateway/internal/adapters/http/handlers"
	"github.com/superdev/ecommerce/gateway/internal/config"
	"github.com/superdev/ecommerce/gateway/internal/domain/service"
)

type getProductHandler struct {
	handlers.AbstractAdminHandler
	productService service.ProductService
}

func NewGetProductHandler(
	productService service.ProductService,
	conf *config.Config,
) handlers.Handler {
	return &getProductHandler{
		productService: productService,
	}
}

func (s *getProductHandler) Pattern() string {
	return "/products/:id"
}

func (s *getProductHandler) Methods() []string {
	return []string{"GET"}
}

func (s *getProductHandler) Handle(c *gin.Context) {
	id := c.Param("id")
	product, err := s.productService.GetProduct(id)
	if err != nil {
		c.JSON(_http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": fmt.Sprintf("Error fetching Product: %s", err.Error()),
		})
		return
	}

	c.JSON(_http.StatusOK, gin.H{
		"status":  "success",
		"message": ("Product fetched successfully"),
		"data":    product,
	})

}
