package authhandlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	auth_dto "github.com/superdev/ecommerce/gateway/internal/adapters/http/dto/auth"
	"github.com/superdev/ecommerce/gateway/internal/adapters/http/handlers"
	"github.com/superdev/ecommerce/gateway/internal/application/service"
	"github.com/superdev/ecommerce/gateway/internal/config"
	"github.com/superdev/ecommerce/gateway/internal/domain/repository"
	"github.com/superdev/ecommerce/gateway/internal/util/auth"
)

type loginAuthhandler struct {
	userRepository  repository.UserRepository
	passwordService service.PasswordService
	conf            *config.Config
}

func NewLoginAuthHandler(
	userRepository repository.UserRepository,
	passwordService service.PasswordService,
	conf *config.Config,
) handlers.Handler {
	return &loginAuthhandler{
		userRepository:  userRepository,
		passwordService: passwordService,
		conf:            conf,
	}
}

func (h *loginAuthhandler) Roles() *[]string {
	// return &[]string{"admin", "user"}
	return nil
}

func (s *loginAuthhandler) Pattern() string {
	return "/login"
}

func (s *loginAuthhandler) Methods() []string {
	return []string{"POST"}
}

func (s *loginAuthhandler) Handle(c *gin.Context) {
	var req auth_dto.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	user, err := s.userRepository.FindByUserName(c, req.Username)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  "error",
			"message": fmt.Sprintf("Invalid username or password"),
		})
		return
	}

	verified, err := s.passwordService.VerifyPassword(req.Password, user.GetHashedPassword())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": fmt.Sprintf("An unkown error occured:", err.Error()),
		})
		return
	}

	if !verified {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  "error",
			"message": fmt.Sprintf("Invalid username or password"),
		})
		return
	}

	token, err := auth.GenerateToken(user.GetId(), s.conf.JWTConfig)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"token":  token,
	})

}
