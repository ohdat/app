package ginmw

import (
	"bytes"
	"fmt"
	"io"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/ohdat/app/middleware/tags/gintags"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

type Config struct {
	TimeFormat   string
	UTC          bool
	RequestBody  bool
	ResponseBody bool
	SkipPaths    []string
}

// Ginzap returns a gin.HandlerFunc using configs
func Ginzap(logger *zap.Logger, conf *Config) gin.HandlerFunc {
	skipPaths := make(map[string]bool, len(conf.SkipPaths))
	for _, path := range conf.SkipPaths {
		skipPaths[path] = true
	}

	return func(c *gin.Context) {
		start := time.Now()
		// some evil middlewares modify this values
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		response := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = response
		bodyBytes, _ := c.GetRawData()
		c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		c.Next()

		if _, ok := skipPaths[path]; !ok {
			end := time.Now()
			latency := end.Sub(start)
			if conf.UTC {
				end = end.UTC()
			}

			fields := []zapcore.Field{
				zap.Int("status", c.Writer.Status()),
				zap.String("method", c.Request.Method),
				zap.String("path", path),
				zap.String("query", query),
				zap.String("ip", c.ClientIP()),
				zap.String("user-agent", c.Request.UserAgent()),
				zap.Duration("latency", latency),
			}
			if conf.TimeFormat != "" {
				fields = append(fields, zap.String("time", end.Format(conf.TimeFormat)))
			}

			// log request ID
			if c.GetHeader("X-Request-Id") != "" {
				fields = append(fields, zap.String("request_id", c.GetHeader("X-Request-Id")))
			}

			// log trace and span ID
			if trace.SpanFromContext(c.Request.Context()).SpanContext().IsValid() {
				traceId := trace.SpanFromContext(c.Request.Context()).SpanContext().TraceID().String()
				c.Set("trace_id", traceId)
				fields = append(fields, zap.String("trace_id", traceId),
					zap.String("span_id", trace.SpanFromContext(c.Request.Context()).SpanContext().SpanID().String()))
			} else {
				traceId := uuid.New().String()
				c.Set("trace_id", traceId)
				fields = append(fields, zap.String("trace_id", traceId))
			}

			// log request body
			if conf.RequestBody && c.Request.Body != nil {
				fields = append(fields, zap.String("request_body", string(bodyBytes)))
			}
			// log response body
			if conf.ResponseBody {
				fields = append(fields, zap.String("response_body", response.body.String()))
			}
			tags := gintags.Values(c)
			for k, v := range tags {
				fields = append(fields, zap.String("tags."+k, fmt.Sprintf("%v", v)))
			}

			if len(c.Errors) > 0 {
				// Append error field if this is an erroneous request.
				for _, e := range c.Errors.Errors() {
					logger.Error(e, fields...)
				}
			} else {
				logger.Info(path, fields...)
			}
		}
	}
}
