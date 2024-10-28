package service

import "github.com/superdev/ecommerce/gateway/internal/domain/entity"

type UserService interface {
	GetUser(id string) (entity.User, error)
	Save(entity.User) (entity.User, error)
	ListUsers() ([]entity.User, error)
	Delete(product entity.User) error
}
