package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/lubosgarancovsky/go-kit"
)

func RoleMiddleware(roles []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userRole := c.Request.Header.Get("X-User-Role")
		if userRole == "" {
			c.Error(go_kit.ErrUnauthorized)
			c.Abort()
			return
		}

		for _, role := range roles {
			if strings.ToLower(role) == strings.ToLower(userRole) {
				c.Next()
				return
			}
		}

		c.Error(go_kit.ErrForbidden)
		c.Abort()
		return
	}
}
