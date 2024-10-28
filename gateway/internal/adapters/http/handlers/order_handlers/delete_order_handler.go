package orderhandlers

import (
	"fmt"
	_http "net/http"

	"github.com/gin-gonic/gin"
	"github.com/superdev/ecommerce/gateway/internal/adapters/http/handlers"
	"github.com/superdev/ecommerce/gateway/internal/config"
	"github.com/superdev/ecommerce/gateway/internal/domain/service"
)

type deleteOrderHandler struct {
	handlers.AbstractUserHandler
	orderService service.OrderService
}

func NewDeleteOrderHandler(
	orderService service.OrderService,
	conf *config.Config,
) handlers.Handler {
	return &deleteOrderHandler{
		orderService: orderService,
	}
}

func (s *deleteOrderHandler) Pattern() string {
	return "/orders/:id"
}

func (s *deleteOrderHandler) Methods() []string {
	return []string{"DELETE"}
}

func (s *deleteOrderHandler) Handle(c *gin.Context) {
	id := c.Param("id")
	order, err := s.orderService.GetOrder(id)
	if err != nil {
		c.JSON(_http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": fmt.Sprintf("Error fetching Order: %s", err.Error()),
		})
		return
	}
	err = s.orderService.Delete(order)
	if err != nil {
		c.JSON(_http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": fmt.Sprintf("Could not delete Order: %s", err.Error()),
		})
		return
	}

	c.JSON(_http.StatusOK, gin.H{
		"status":  "success",
		"message": ("Order deleted successfully"),
		"data":    nil, //order,
	})

}
