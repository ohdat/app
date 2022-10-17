package logger

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/ohdat/app/utils"
	"go.uber.org/zap"
	"sync"
)

type Logger struct {
	*zap.Logger
}

func (s *Logger) Println(v ...interface{}) {
	s.Info(fmt.Sprintln(v...))
}

func (s *Logger) Printf(format string, v ...interface{}) {
	s.Info(fmt.Sprintf(format, v...))
}

func (s *Logger) ErrPrintln(v ...interface{}) {
	s.Error(fmt.Sprintln(v...))
}

func (s *Logger) Trace(id string) *zap.Logger {
	if id == "" {
		id = uuid.New().String()
	}
	return s.With(zap.String("trace_id", id))
}

var zapLogger *Logger

var once sync.Once

func Get(app string) *Logger {
	once.Do(func() {
		zapLogger = &Logger{
			NewZapLogger(app),
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
