package service

import "github.com/superdev/ecommerce/gateway/internal/domain/entity"

type OrderService interface {
	// CheckOrderAvailability(id string) (*bool, error)
	Save(entity.Order) (entity.Order, error)
	GetOrder(id string) (entity.Order, error)
	ListOrders() ([]entity.Order, error)
	Delete(order entity.Order) error
}
