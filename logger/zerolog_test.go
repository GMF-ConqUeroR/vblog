package logger_test

import (
	"testing"

	"gitee.com/xpigpig/vblog/logger"
)

func TestLogger(t *testing.T) {
	logger.Log().Debug().Msg("This is  test!")
}
