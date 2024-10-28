package service

import (
	"context"
	"fmt"

	"github.com/superdev/ecommerce/gateway/internal/domain/entity"
	"github.com/superdev/ecommerce/gateway/internal/domain/repository"
	"github.com/superdev/ecommerce/gateway/internal/domain/service"
	"go.uber.org/zap"
)

func NewProductService(
	repo repository.ProductRepository,
	logger zap.Logger,
) (service.ProductService, error) {

	return &productService{
		repo:   repo,
		logger: logger.With(zap.Namespace("app.service.product")),
	}, nil
}

type productService struct {
	repo   repository.ProductRepository
	logger *zap.Logger
}

func (s *productService) CheckProductAvailability(id string, quantity int64) (*bool, error) {
	return s.repo.CheckProductAvailability(context.Background(), id, quantity)
}

func (s *productService) GetProduct(id string) (entity.Product, error) {
	s.logger.Debug(fmt.Sprintf("Fetching product with id \"%s\"", id))
	//
	product, err := s.repo.FindById(context.Background(), id)
	//
	if err != nil {
		s.logger.Error(fmt.Sprintf("Fetching product failed id \"%s\"", err.Error()))
	} else {
		s.logger.Error(fmt.Sprintf("Fetching product succeded id \"%v\"", product))
	}
	return product, err
}
func (s *productService) ListProducts() ([]entity.Product, error) {
	return s.repo.FindAll(context.Background())
}

func (s *productService) Save(product entity.Product) (entity.Product, error) {
	if product.GetId() != "" && product.GetId() != "0" {
		err := s.repo.Update(context.Background(), product)
		return product, err
	}
	return s.repo.Create(context.Background(), product)
}

func (s *productService) Delete(product entity.Product) error {
	return s.repo.Delete(context.Background(), product)
}
