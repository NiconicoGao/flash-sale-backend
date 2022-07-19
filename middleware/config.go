package middleware

import (
	"flash-sale-backend/utils"

	"github.com/gin-gonic/gin"
)

func ConfigMiddleware(config *utils.MyConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Set example variable
		c.Set("config", config)
		c.Next()
	}
}
