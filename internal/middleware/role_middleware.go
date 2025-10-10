package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/lubosgarancovsky/go-kit/api_err"
)

func RoleMiddleware(roles []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userRole := c.Request.Header.Get("X-User-Role")
		if userRole == "" {
			c.Error(api_err.ErrUnauthorized)
			c.Abort()
			return
		}

		for _, role := range roles {
			if strings.ToLower(role) == strings.ToLower(userRole) {
				c.Next()
				return
			}
		}

		c.Error(api_err.ErrForbidden)
		c.Abort()
		return
	}
}
