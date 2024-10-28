package service

import (
	"context"
	"fmt"

	"github.com/superdev/ecommerce/gateway/internal/domain/entity"
	"github.com/superdev/ecommerce/gateway/internal/domain/repository"
	"github.com/superdev/ecommerce/gateway/internal/domain/service"
	"go.uber.org/zap"
)

func NewOrderService(
	repo repository.OrderRepository,
	logger zap.Logger,
) (service.OrderService, error) {

	return &orderService{
		repo:   repo,
		logger: logger.With(zap.Namespace("app.service.order")),
	}, nil
}

type orderService struct {
	repo   repository.OrderRepository
	logger *zap.Logger
}

func (s *orderService) GetOrder(id string) (entity.Order, error) {
	s.logger.Debug(fmt.Sprintf("Fetching order with id \"%s\"", id))
	//
	order, err := s.repo.FindById(context.Background(), id)
	//
	if err != nil {
		s.logger.Error(fmt.Sprintf("Fetching order failed id \"%s\"", err.Error()))
	} else {
		s.logger.Error(fmt.Sprintf("Fetching order succeded id \"%v\"", order))
	}
	return order, err
}
func (s *orderService) ListOrders() ([]entity.Order, error) {
	return s.repo.FindAll(context.Background())
}

func (s *orderService) Save(order entity.Order) (entity.Order, error) {
	if order.GetId() != "" && order.GetId() != "0" {
		err := s.repo.Update(context.Background(), order)
		return order, err
	}
	return s.repo.Create(context.Background(), order)
}

func (s *orderService) Delete(order entity.Order) error {
	return s.repo.Delete(context.Background(), order)
}
