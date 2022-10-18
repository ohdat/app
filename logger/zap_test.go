package logger

import (
	"testing"
)

func TestZap(t *testing.T) {
	var logger = Get("test")
	logger.Info("testInfo")
	logger.Warn("testWarn")
}
