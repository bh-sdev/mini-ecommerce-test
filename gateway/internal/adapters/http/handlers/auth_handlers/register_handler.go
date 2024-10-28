package authhandlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	auth_dto "github.com/superdev/ecommerce/gateway/internal/adapters/http/dto/auth"
	"github.com/superdev/ecommerce/gateway/internal/adapters/http/handlers"

	"github.com/superdev/ecommerce/gateway/internal/application/service"
	"github.com/superdev/ecommerce/gateway/internal/config"
	"github.com/superdev/ecommerce/gateway/internal/domain/entity"
	"github.com/superdev/ecommerce/gateway/internal/domain/repository"
	_service "github.com/superdev/ecommerce/gateway/internal/domain/service"
)

type registerAuthhandler struct {
	userRepository  repository.UserRepository
	passwordService service.PasswordService
	userservice     _service.UserService
	conf            *config.Config
}

func NewRegisterAuthHandler(
	userservice _service.UserService,
	userRepository repository.UserRepository,
	passwordService service.PasswordService,
	conf *config.Config,
) handlers.Handler {
	return &registerAuthhandler{
		userservice:     userservice,
		userRepository:  userRepository,
		passwordService: passwordService,
		conf:            conf,
	}
}

func (h *registerAuthhandler) Roles() *[]string {
	// return &[]string{"admin", "user"}
	return nil
}

func (s *registerAuthhandler) Pattern() string {
	return "/register"
}

func (s *registerAuthhandler) Methods() []string {
	return []string{"POST"}
}

func (s *registerAuthhandler) Handle(c *gin.Context) {
	var req auth_dto.RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	hashedPassword, err := s.passwordService.HashPassword(req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": fmt.Sprintf("An unkown error occured:", err.Error()),
		})
		return
	}

	user := entity.NewEmptyUser()
	user.SetUserName(req.Username)
	user.SetHashedPassword(hashedPassword)

	_user, err := s.userservice.Save(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": fmt.Sprintf("Error Creating account: %s", err.Error()),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   _user,
	})

}
