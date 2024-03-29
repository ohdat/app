package logger

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/ohdat/app/utils"
	"go.uber.org/zap"
	"sync"
)

type Logger struct {
	TraceID string
	*zap.Logger
}

func (s *Logger) Info(msg string, fields ...zap.Field) {
	fields = append(fields, zap.String("trace_id", s.TraceID))
	s.WithOptions(zap.AddCallerSkip(1)).Info(msg, fields...)
}

func (s *Logger) Warn(msg string, fields ...zap.Field) {
	fields = append(fields, zap.String("trace_id", s.TraceID))
	s.WithOptions(zap.AddCallerSkip(1)).Warn(msg, fields...)
}

func (s *Logger) Error(msg string, fields ...zap.Field) {
	fields = append(fields, zap.String("trace_id", s.TraceID))
	s.WithOptions(zap.AddCallerSkip(1)).Error(msg, fields...)
}

func (s *Logger) Debug(msg string, fields ...zap.Field) {
	fields = append(fields, zap.String("trace_id", s.TraceID))
	s.WithOptions(zap.AddCallerSkip(1)).Debug(msg, fields...)
}
func (s *Logger) Panic(msg string, fields ...zap.Field) {
	fields = append(fields, zap.String("trace_id", s.TraceID))
	s.WithOptions(zap.AddCallerSkip(1)).Panic(msg, fields...)
}

func (s *Logger) Fatal(msg string, fields ...zap.Field) {
	fields = append(fields, zap.String("trace_id", s.TraceID))
	s.WithOptions(zap.AddCallerSkip(1)).Fatal(msg, fields...)
}

func (s *Logger) Println(v ...interface{}) {
	s.WithOptions(zap.AddCallerSkip(1)).Info(fmt.Sprintln(v...), zap.String("trace_id", s.TraceID))
}

func (s *Logger) Printf(format string, v ...interface{}) {
	s.WithOptions(zap.AddCallerSkip(1)).Info(fmt.Sprintf(format, v...), zap.String("trace_id", s.TraceID))
}

func (s *Logger) ErrPrintln(v ...interface{}) {
	s.WithOptions(zap.AddCallerSkip(1)).Error(fmt.Sprintln(v...), zap.String("trace_id", s.TraceID))
}

func (s *Logger) Trace(id string) {
	if id == "" {
		id = uuid.New().String()
	}
	s.TraceID = id
}
func (s *Logger) Copy() *Logger {
	copy := *s
	return &copy
}

var zapLogger *Logger

var once sync.Once

func Get(app string) *Logger {
	once.Do(func() {
		zapLogger = &Logger{
			Logger: NewZapLogger(app),
		}
	})
	return zapLogger
}

func NewZapLogger(app string) *zap.Logger {
	var logger *zap.Logger
	if utils.IsDev() {
		logger, _ = zap.NewDevelopment()
	} else {
		logger, _ = zap.NewProduction()
	}
	return logger.With(zap.String("app", app))
}
