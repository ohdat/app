package ginm

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

func TraceMiddleware(zapLogger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		//开始时间
		var traceId = c.GetHeader("X-B3-Traceid")
		if traceId == "" {
			traceId = uuid.New().String()
		}
		var requestId = c.GetHeader("X-Request-Id")
		zapLogger.With(zap.String("traceId", traceId), zap.String("requestId", requestId))
		c.Next()
	}
}
