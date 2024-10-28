package repository

import "github.com/superdev/ecommerce/gateway/internal/domain/entity"

type OrderRepository interface {
	Repository[entity.Order]
}
