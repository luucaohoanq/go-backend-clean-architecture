package middleware

import (
	"log/slog"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func LoggingMiddleware(baseLogger *slog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		reqLogger := baseLogger.With(
			"requestId", uuid.NewString(),
			"path", c.Request.URL.Path,
		)

		c.Set("logger", reqLogger)

		c.Next()

		reqLogger.Info("Request handled",
			"method", c.Request.Method,
			"status", c.Writer.Status(),
			"duration", time.Since(start),
		)
	}
}
