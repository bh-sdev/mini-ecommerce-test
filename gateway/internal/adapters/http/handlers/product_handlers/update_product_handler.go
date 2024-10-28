package producthandlers

import (
	"fmt"
	_http "net/http"

	"github.com/gin-gonic/gin"
	"github.com/superdev/ecommerce/gateway/internal/adapters/http/dto/product"
	"github.com/superdev/ecommerce/gateway/internal/adapters/http/handlers"
	"github.com/superdev/ecommerce/gateway/internal/config"
	"github.com/superdev/ecommerce/gateway/internal/domain/entity"
	"github.com/superdev/ecommerce/gateway/internal/domain/service"
)

type updateProductHandler struct {
	handlers.AbstractAdminHandler
	productService service.ProductService
}

func NewUpdateProductHandler(
	productService service.ProductService,
	conf *config.Config,
) handlers.Handler {
	return &updateProductHandler{
		productService: productService,
	}
}

func (s *updateProductHandler) Pattern() string {
	return "/products/:id"
}

func (s *updateProductHandler) Methods() []string {
	return []string{"PATCH", "PUT"}
}

func (s *updateProductHandler) Handle(c *gin.Context) {
	var req product.UpdateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(_http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	id := c.Param("id")
	_product, err := s.productService.GetProduct(id)
	if err != nil {
		c.JSON(_http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": fmt.Sprintf("Error fetching Product: %s", err.Error()),
		})
		return
	}

	product := entity.NewProductFromEntity(_product)

	if nil != req.Name {
		product.SetName(*req.Name)
	}
	if nil != req.Description {
		product.SetDescription(*req.Description)
	}
	if nil != req.AvailableQuantity {
		product.SetQuantityAvailable(*req.AvailableQuantity)
	}
	if nil != req.Price {
		product.SetPrice(int64(*req.Price))
	}

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
