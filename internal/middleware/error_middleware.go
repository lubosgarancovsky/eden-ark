package middleware

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/lubosgarancovsky/go-kit"
	"gorm.io/gorm"
)

var serviceID = "eden-ark"

func ErrorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			lastErr := c.Errors.Last().Err
			correlationID, err := uuid.Parse(c.GetHeader("correlationId"))
			if err != nil {
				correlationID = uuid.Nil
			}

			var apiErr *go_kit.ApiError
			if errors.As(lastErr, &apiErr) {
				apiErr.Log()
				c.JSON(apiErr.HTTPStatus, apiErr.ToJSON(serviceID, correlationID.String()))
				return
			}

			if errors.Is(lastErr, gorm.ErrRecordNotFound) {
				apiErr := go_kit.ErrNotFound
				apiErr.Log()
				c.JSON(apiErr.HTTPStatus, apiErr.ToJSON(serviceID, correlationID.String()))
				return
			}

			unknownError := go_kit.Unknown(lastErr)
			unknownError.Log()
			c.JSON(http.StatusInternalServerError, unknownError.ToJSON(serviceID, correlationID.String()))
		}
	}
}
