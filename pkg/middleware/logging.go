package middleware

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// Log to Console
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Processing request
		c.Next()
		// Request mode
		reqMethod := c.Request.Method
		// Request routing
		reqUri := c.Request.RequestURI
		// Status code
		statusCode := c.Writer.Status()
		// Request IP
		clientIP := c.ClientIP()
		// Log format
		log.SetFormatter(&log.TextFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
		})
		log.SetFormatter(&log.JSONFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
		})

		log.Infof("| %3d | %15s | %s | %s |",
			statusCode,
			clientIP,
			reqMethod,
			reqUri,
		)

	}

}
