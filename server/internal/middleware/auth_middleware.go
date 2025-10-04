package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserContext struct {
	ID   uuid.UUID
	Role string
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("user", &UserContext{ID: uuid.Nil, Role: "ADMIN"})
		c.Next()
	}
}
