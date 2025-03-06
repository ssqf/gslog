package gslog

import (
	"fmt"
	"io"
	"log"
	"os"
)

type LogLevel int

const (
	TRACE LogLevel = iota // 跟踪级别
	DEBUG                 // 调试级别
	INFO                  // 信息级别
	WARN                  // 警告级别
	ERROR                 // 错误级别
	NONE                  // 关闭日志
)

var logLevel = INFO
var logWriter io.Writer
var logger *log.Logger

// 获取纯文本级别字符串
func getTextLevel(level LogLevel) string {
	return []string{"[TRACE] ", "[DEBUG] ", "[INFO] ", "[WARN] ", "[ERROR] ", "[NONE] "}[level]
}
func init() {
	logWriter = os.Stdout
	logger = log.New(logWriter, "", log.Lshortfile|log.LstdFlags)
}

func logout(level LogLevel, format string, v ...any) {
	if logLevel < level {
		return
	}

	logger.Output(3, getTextLevel(level)+fmt.Sprintf(format, v...))
}
func Info(format string, v ...any) {
	logout(INFO, format, v...)
}

func Warn(format string, v ...any) {
	logout(WARN, format, v...)
}

func Debug(format string, v ...any) {
	logout(DEBUG, format, v...)
}

func Error(format string, v ...any) {
	logout(ERROR, format, v...)
}

func Trace(format string, v ...any) {
	logout(TRACE, format, v...)
}

func SetOutput(w io.Writer) {
	logWriter = w
	logger.SetOutput(w)
}

func SetLevel(level LogLevel) {
	logLevel = level
}

func GetLevel() LogLevel {
	return logLevel
}

func SetFlags(flags int) {
	logger.SetFlags(flags)
}

func Fatalf(format string, v ...any) {
	logger.Output(2, fmt.Sprintf("[FATAL] "+format, v...))
	os.Exit(1)
}

func Printf(format string, v ...any) {
	if logLevel == NONE {
		return
	}
	logger.Output(2, fmt.Sprintf(format, v...))
}

func Println(v ...any) {
	if logLevel == NONE {
		return
	}

	logger.Output(2, fmt.Sprintln(v...))
}

func GetDefaultLogger() *log.Logger {
	return logger
}
