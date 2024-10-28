package repository

import "context"

type Repository[T any] interface {
	Create(ctx context.Context, entity T) (T, error)
	FindById(ctx context.Context, id string) (T, error)
	FindAll(ctx context.Context) ([]T, error)
	Update(ctx context.Context, entity T) error
	Delete(ctx context.Context, entity T) error
}
