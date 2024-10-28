package service

import (
	"context"
	"fmt"

	"github.com/superdev/ecommerce/gateway/internal/domain/entity"
	"github.com/superdev/ecommerce/gateway/internal/domain/repository"
	"github.com/superdev/ecommerce/gateway/internal/domain/service"
	"go.uber.org/zap"
)

func NewUserService(
	repo repository.UserRepository,
	passwordVerifier PasswordService,
	logger zap.Logger,
) (service.UserService, error) {

	return &userService{
		repo:             repo,
		passwordVerifier: passwordVerifier,
		logger:           logger.With(zap.Namespace("app.service.user")),
	}, nil
}

type userService struct {
	repo             repository.UserRepository
	passwordVerifier PasswordService
	logger           *zap.Logger
}

func (s *userService) GetUser(id string) (entity.User, error) {
	s.logger.Debug(fmt.Sprintf("Fetching user with id \"%s\"", id))
	//
	user, err := s.repo.FindById(context.Background(), id)
	//
	if err != nil {
		s.logger.Error(fmt.Sprintf("Fetching user failed id \"%s\"", err.Error()))
	} else {
		s.logger.Error(fmt.Sprintf("Fetching user succeded id \"%v\"", user))
	}
	return user, err
}
func (s *userService) ListUsers() ([]entity.User, error) {
	return s.repo.FindAll(context.Background())
}

func (s *userService) Save(user entity.User) (entity.User, error) {
	if password := user.GetPassword(); password != "" {
		hashedPassword, err := s.passwordVerifier.HashPassword(password)
		if err != nil {
			return nil, err
		}
		user.SetHashedPassword(hashedPassword)
	}
	if user.GetId() != "" && user.GetId() != "0" {
		err := s.repo.Update(context.Background(), user)
		return user, err
	}
	return s.repo.Create(context.Background(), user)
}

func (s *userService) Delete(user entity.User) error {
	return s.repo.Delete(context.Background(), user)
}
