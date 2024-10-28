package producthandlers

import (
	"fmt"
	_http "net/http"

	"github.com/gin-gonic/gin"
	// "github.com/superdev/ecommerce/gateway/internal/adapters/http"
	"github.com/superdev/ecommerce/gateway/internal/adapters/http/dto/product"
	"github.com/superdev/ecommerce/gateway/internal/adapters/http/handlers"
	"github.com/superdev/ecommerce/gateway/internal/config"
	"github.com/superdev/ecommerce/gateway/internal/domain/entity"
	"github.com/superdev/ecommerce/gateway/internal/domain/service"
)

type createProductHandler struct {
	handlers.AbstractAdminHandler
	productService service.ProductService
}

func NewCreateProductHandler(
	productService service.ProductService,
	conf *config.Config,
) handlers.Handler {
	return &createProductHandler{
		productService: productService,
	}
}

func (s *createProductHandler) Pattern() string {
	return "/products"
}

func (s *createProductHandler) Methods() []string {
	return []string{"POST"}
}

func (s *createProductHandler) Handle(c *gin.Context) {
	var req product.CreateProductRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(_http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}
	product := entity.NewEmptyProduct()

	product.SetName(req.Name)
	product.SetDescription(req.Description)
	product.SetQuantityAvailable(req.AvailableQuantity)
	product.SetPrice(int64(req.Price))

	res, err := s.productService.Save(product)
	if err != nil {
		c.JSON(_http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": fmt.Sprintf("Error saving Product: %s", err.Error()),
		})
		return
	}
	c.JSON(_http.StatusOK, gin.H{
		"status":  "success",
		"message": ("Product saved successfully"),
		"data":    res,
	})
}
