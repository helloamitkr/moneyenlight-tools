package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/helloamitkr/moneyenlight-tools/pkg/logger"
)

// RequestLogger logs structured JSON for every HTTP request
func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()

		latency := time.Since(start)
		method := c.Request.Method
		path := c.Request.URL.Path
		status := c.Writer.Status()
		clientIP := c.ClientIP()

		logger.Log.WithFields(map[string]interface{}{
			"method":  method,
			"path":    path,
			"status":  status,
			"latency": latency.String(),
			"ip":      clientIP,
		}).Info("HTTP request")
	}
}
