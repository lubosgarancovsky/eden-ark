package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/lubosgarancovsky/go-kit"
	"gorm.io/gorm"
)

func AppStateValidationMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		if db == nil {
			c.Error(go_kit.ErrInternalServer.WithMessage("No database connection"))
			c.Abort()
		}

		c.Next()
	}
}
