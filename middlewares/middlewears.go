package middlewares

import (
	"github.com/gin-gonic/gin"
	"go-social-network/server/utils"
	"net/http"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := utils.TokenIsValid(c); err != nil {
			c.String(http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}
		c.Next()
	}
}
