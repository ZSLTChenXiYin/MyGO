package middleware

import (
	"time"

	"github.com/ZSLTChenXiYin/MyGO/logger"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func StdLogger(logger *logger.StdLogger) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		request_uuid := c.GetHeader("Request-Uuid")

		if len(c.Errors) > 0 {
			logger.Errorf("ID: %s, Status: %d, Method: %s, Path: %s, Query: %v, Error: %s", request_uuid, c.Writer.Status(), c.Request.Method, c.Request.URL.Path, c.Request.URL.RawQuery, c.Errors.String())
		} else {
			logger.Infof("ID: %s, Status: %d, Method: %s, Path: %s, Query: %v", request_uuid, c.Writer.Status(), c.Request.Method, c.Request.URL.Path, c.Request.URL.RawQuery)
		}
	}
}

func ZapLogger(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		request_uuid := c.GetHeader("Request-Uuid")

		// 处理请求
		c.Next()

		end := time.Now()
		latency := end.Sub(start)

		if len(c.Errors) > 0 {
			// 如果有错误，记录错误信息
			logger.Error("HTTP Request",
				zap.Int("status", c.Writer.Status()),
				zap.String("method", c.Request.Method),
				zap.String("path", path),
				zap.String("query", query),
				zap.String("ip", c.ClientIP()),
				zap.String("user-agent", c.Request.UserAgent()),
				zap.Duration("latency", latency),
				zap.String("time", end.Format(time.RFC3339)),
				zap.String("request_uuid", request_uuid),
				zap.String("errors", c.Errors.String()),
			)
		} else {
			// 记录正常的请求日志
			logger.Info("HTTP Request",
				zap.Int("status", c.Writer.Status()),
				zap.String("method", c.Request.Method),
				zap.String("path", path),
				zap.String("query", query),
				zap.String("ip", c.ClientIP()),
				zap.String("user-agent", c.Request.UserAgent()),
				zap.Duration("latency", latency),
				zap.String("time", end.Format(time.RFC3339)),
				zap.String("request_uuid", request_uuid),
			)
		}
	}
}
