package repository

import (
	"context"

	"github.com/superdev/ecommerce/gateway/internal/domain/entity"
)

type UserRepository interface {
	Repository[entity.User]
	FindByUserName(ctx context.Context, Username string) (entity.User, error)
}
