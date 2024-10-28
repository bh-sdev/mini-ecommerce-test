package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/superdev/ecommerce/gateway/internal/adapters/http/handlers"
	"github.com/superdev/ecommerce/gateway/internal/domain/entity"
	"github.com/superdev/ecommerce/gateway/internal/util/str"
)

func Authenticate(h handlers.Handler) gin.HandlerFunc {
	roles := h.Roles()
	return func(c *gin.Context) {
		if nil == roles || len(*roles) == 0 {
			c.Next()
			return
		} else {
			_user, exists := c.Get("user")
			if !exists {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Account not found"})
				c.Abort()
				return
			}
			user, _ := _user.(entity.User)
			if !str.StringInSlice(user.GetRole(), *roles) {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
				c.Abort()
				return
			}
		}
		c.Next()
	}
}
