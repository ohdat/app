package logger

import (
	"testing"
	"time"
)

func TestZap(t *testing.T) {
	var logger = NewZapLogger("test")
	logger.Info("testInfo")
}

func TestTrace(t *testing.T) {
	var logger = Get("test")
	go func(logger Logger) {
		logger.Logger = logger.Trace("22")
		time.Sleep(time.Second)
		logger.Println("111111111")
	}(*logger)
	time.Sleep(time.Nanosecond * 10)

	go func(logger Logger) {
		logger.Logger = logger.Trace("33")
		time.Sleep(time.Second)
		logger.Println("2222222")
	}(*logger)

	time.Sleep(time.Second * 2)
}
