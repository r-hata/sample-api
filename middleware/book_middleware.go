package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/r-hata/sample-api/logger"
	"go.uber.org/zap"
)

func RecordUaAndTime(c *gin.Context) {
	startTime := time.Now()
	ua := c.GetHeader("User-Agent")
	c.Next()
	logger.Info(
		"incoming request",
		zap.String("path", c.Request.URL.Path),
		zap.String("ua", ua),
		zap.Int("status", c.Writer.Status()),
		zap.Duration("elapsed", time.Now().Sub(startTime)),
	)
}
