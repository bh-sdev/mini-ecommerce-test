package orderhandlers

import (
	"fmt"
	_http "net/http"

	"github.com/gin-gonic/gin"
	"github.com/superdev/ecommerce/gateway/internal/adapters/http/dto/order"
	"github.com/superdev/ecommerce/gateway/internal/adapters/http/handlers"
	"github.com/superdev/ecommerce/gateway/internal/config"
	"github.com/superdev/ecommerce/gateway/internal/domain/entity"
	"github.com/superdev/ecommerce/gateway/internal/domain/service"
)

type createOrderHandler struct {
	handlers.AbstractUserHandler
	productService service.ProductService
	orderService   service.OrderService
}

func NewCreateOrderHandler(
	productService service.ProductService,
	orderService service.OrderService,
	conf *config.Config,
) handlers.Handler {
	return &createOrderHandler{
		orderService:   orderService,
		productService: productService,
	}
}

func (s *createOrderHandler) Pattern() string {
	return "/orders"
}

func (s *createOrderHandler) Methods() []string {
	return []string{"POST"}
}

func (s *createOrderHandler) Handle(c *gin.Context) {
	var req order.CreateOrderRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(_http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}
	order := entity.NewEmptyOrder()
	for _, iItem := range req.Items {
		product, err := s.productService.GetProduct(iItem.ProductId)
		if err != nil {
			c.JSON(_http.StatusBadRequest, gin.H{
				"status":  "error",
				"message": fmt.Sprintf("Error fetching product with id[%s]: %s", iItem.ProductId, err.Error()),
			})
			return
		}
		var item = entity.NewEmptyOrderItem()
		item.SetProductId(iItem.ProductId)
		item.SetPrice(product.GetPrice())
		item.SetQuantity(iItem.Quantity)
		//
		order.AddItem(item)
	}

	res, err := s.orderService.Save(order)
	if err != nil {
		c.JSON(_http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": fmt.Sprintf("Error saving Order: %s", err.Error()),
		})
		return
	}
	c.JSON(_http.StatusOK, gin.H{
		"status":  "success",
		"message": ("Order saved successfully"),
		"data":    res,
	})
}
