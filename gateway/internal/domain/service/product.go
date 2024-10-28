package service

import "github.com/superdev/ecommerce/gateway/internal/domain/entity"

type ProductService interface {
	CheckProductAvailability(id string, quantity int64) (*bool, error)
	GetProduct(id string) (entity.Product, error)
	Save(entity.Product) (entity.Product, error)
	ListProducts() ([]entity.Product, error)
	Delete(product entity.Product) error
}
