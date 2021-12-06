package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

// Log to Console
func Logger() gin.HandlerFunc {
	logger := logrus.New()
	//Set output
	logger.Out = os.Stdout
	//Set log level
	logger.SetLevel(logrus.DebugLevel)
	//Format log
	logger.SetFormatter(&logrus.TextFormatter{})
	return func(c *gin.Context) {
		// start time
		startTime := time.Now()
		// Processing request
		c.Next()
		// End time
		endTime := time.Now()
		// execution time
		latencyTime := endTime.Sub(startTime)
		// Request mode
		reqMethod := c.Request.Method
		// Request routing
		reqUri := c.Request.RequestURI
		// Status code
		statusCode := c.Writer.Status()
		// Request IP
		clientIP := c.ClientIP()
		// Log format
		logger.SetFormatter(&logrus.TextFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
		})
		logger.SetFormatter(&logrus.JSONFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
		})

		logger.Infof("| %3d | %13v | %15s | %s | %s |",
			statusCode,
			latencyTime,
			clientIP,
			reqMethod,
			reqUri,
		)

	}

}
