package middleware

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lubosgarancovsky/eden-arc/pkg/errors"
)

var serviceID = "eden-arc"

func ErrorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			lastErr := c.Errors.Last().Err
			correlationID := c.GetHeader("correlationId")
			currentTime := time.Now().Format(time.RFC3339)

			if apiErr, ok := lastErr.(*errors.APIError); ok {
				log.Printf("[ERROR] %d %s \"%s\": %v", apiErr.HTTPStatus, apiErr.Code, apiErr.Message, apiErr.Error())
				c.JSON(apiErr.HTTPStatus, gin.H{
					"code":          apiErr.Code,
					"message":       apiErr.Message,
					"correlationId": correlationID,
					"serviceId":     serviceID,
					"timestamp":     currentTime,
				})
				return
			}

			log.Printf("[ERROR] 500 INTERNAL_SERVER_ERROR \"Internal server error\": %v", lastErr)
			c.JSON(http.StatusInternalServerError, gin.H{
				"code":          "INTERNAL_SERVER_ERROR",
				"message":       "Internal server error",
				"correlationId": correlationID,
				"serviceId":     serviceID,
				"timestamp":     currentTime,
			})
		}
	}
}
