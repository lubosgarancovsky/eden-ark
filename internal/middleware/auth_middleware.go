package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-ark/internal/model"
	"github.com/lubosgarancovsky/go-kit"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.Request.Header.Get("X-User-Id")
		if userID == "" {
			c.Error(go_kit.ErrUnauthorized)
			c.Abort()
			return
		}

		userRole := c.Request.Header.Get("X-User-Role")
		if userRole == "" {
			c.Error(go_kit.ErrUnauthorized)
			c.Abort()
			return
		}

		userUUID, err := uuid.Parse(userID)
		if err != nil {
			c.Error(go_kit.ErrUnauthorized)
			c.Abort()
			return
		}

		c.Set("user", &model.UserContext{ID: userUUID, Role: userRole})
		c.Next()
	}
}
