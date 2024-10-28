package orderhandlers

import (
	"fmt"
	_http "net/http"

	"github.com/gin-gonic/gin"
	"github.com/superdev/ecommerce/gateway/internal/adapters/http/handlers"
	"github.com/superdev/ecommerce/gateway/internal/config"
	"github.com/superdev/ecommerce/gateway/internal/domain/service"
)

type getOrderHandler struct {
	handlers.AbstractUserHandler
	orderService service.OrderService
}

func NewGetOrderHandler(
	orderService service.OrderService,
	conf *config.Config,
) handlers.Handler {
	return &getOrderHandler{
		orderService: orderService,
	}
}

func (s *getOrderHandler) Pattern() string {
	return "/orders/:id"
}

func (s *getOrderHandler) Methods() []string {
	return []string{"GET"}
}

func (s *getOrderHandler) Handle(c *gin.Context) {
	id := c.Param("id")
	order, err := s.orderService.GetOrder(id)
	if err != nil {
		c.JSON(_http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": fmt.Sprintf("Error fetching Order: %s", err.Error()),
		})
		return
	}

	c.JSON(_http.StatusOK, gin.H{
		"status":  "success",
		"message": ("Order fetched successfully"),
		"data":    order,
	})

}
