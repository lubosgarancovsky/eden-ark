package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/lubosgarancovsky/go-kit/api_err"
	"gorm.io/gorm"
)

func AppStateValidationMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		if db == nil {
			c.Error(api_err.ErrInternalServer.WithMessage("No database connection"))
			c.Abort()
		}

		c.Next()
	}
}
