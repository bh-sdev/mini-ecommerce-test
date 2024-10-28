package gormrepo

import (
	"context"

	"github.com/superdev/ecommerce/gateway/internal/domain/entity"
	"github.com/superdev/ecommerce/gateway/internal/domain/repository"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type userRepository struct {
	db     *gorm.DB
	logger zap.Logger
}

func NewGormUserRepository(
	db *gorm.DB,
	logger zap.Logger,
) repository.UserRepository {
	return &userRepository{
		db:     db,
		logger: logger,
	}
}

func (s *userRepository) Create(ctx context.Context, entity entity.User) (entity.User, error) {
	dUser := fromUserEntity(entity)
	if err := s.db.Save(dUser).Error; err != nil {
		return nil, err
	}
	return dUser, nil
}
func (s *userRepository) Delete(ctx context.Context, entity entity.User) error {
	dUser := fromUserEntity(entity)
	return s.db.Delete(dUser).Error
}
func (s *userRepository) FindAll(ctx context.Context) ([]entity.User, error) {
	var _users []GormUser
	if err := s.db.Find(&_users).Error; err != nil {
		return nil, err
	}
	var users = make([]entity.User, len(_users), cap(_users))
	for _, user := range _users {
		users = append(users, &user)
	}
	return users, nil

}

func (s *userRepository) FindById(ctx context.Context, id string) (entity.User, error) {
	var user GormUser
	if err := s.db.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *userRepository) FindByUserName(ctx context.Context, username string) (entity.User, error) {
	var user GormUser
	if err := s.db.Where(" username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *userRepository) Update(ctx context.Context, entity entity.User) error {
	dUser := fromUserEntity(entity)
	return s.db.Save(dUser).Error
}
