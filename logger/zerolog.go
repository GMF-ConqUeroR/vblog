package logger

import (
	"fmt"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/rs/zerolog"
)

var logger *zerolog.Logger

func Log() *zerolog.Logger {
	return logger
}

func initLoggerConfig() {
	output := zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: time.RFC3339,
	}
	output.FormatLevel = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf("| %-6s|", i))
	}
	output.FormatMessage = func(i interface{}) string {
		return fmt.Sprintf("***%s****", i)
	}
	output.FormatFieldName = func(i interface{}) string {
		return fmt.Sprintf("%s:", i)
	}
	output.FormatFieldValue = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf("%s", i))
	}
	log := zerolog.New(output).With().Timestamp().Caller().Logger()
	logger = &log
}

// 1. 每次Import Logger这个包，都要执行init函数
// 2. logger对象不需要重复初始化, 重复初始化可能出问题
// 3. 使用sync once, 无论这个包被导入多少次，initLogger函数只执行一次
var once sync.Once

func init() {
	once.Do(initLoggerConfig)
}
