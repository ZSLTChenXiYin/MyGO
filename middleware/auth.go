package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Auth(key string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authorization := c.GetHeader("Authorization")
		if authorization == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		authorization_slice := strings.Split(authorization, " ")
		if len(authorization_slice) != 2 || authorization_slice[0] != "Bearer" {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		c.Set(key, authorization_slice[1])

		c.Next()
	}
}
