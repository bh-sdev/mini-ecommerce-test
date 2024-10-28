package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/superdev/ecommerce/gateway/internal/config"
)

type Claims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

// GenerateToken creates a JWT token for a given user ID
func GenerateToken(userID string, config config.JWTConfig) (string, error) {
	claims := Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(config.ExpiresIn)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.SigningKey))
}
