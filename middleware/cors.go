package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Cors(allow_headers string, expose_headers string) gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")
		if origin != "" {
			c.Header("Access-Control-Allow-Origin", origin)
		} else {
			c.Header("Access-Control-Allow-Origin", "*")
		}

		if allow_headers == "" {
			allow_headers = "Content-Type, Authorization, Request-Uuid"
		}
		if expose_headers == "" {
			expose_headers = "Content-Length, Content-Type"
		}

		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", allow_headers)
		c.Header("Access-Control-Expose-Headers", expose_headers) // 可选：允许前端访问的响应头
		c.Header("Access-Control-Allow-Credentials", "true")      // 可选：如果需要跨域携带 Cookie

		// 处理预检请求
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}

		// 继续执行后续中间件和路由处理
		c.Next()
	}
}
