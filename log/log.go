/*
Package log provides support for logging to stdout and stderr.

Log entries will be logged in the following format:

	timestamp hostname tag[pid]: SEVERITY Message
*/
package log

import (
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/natefinch/lumberjack"
	log "github.com/sirupsen/logrus"
)

type ConfdFormatter struct {
}

func (c *ConfdFormatter) Format(entry *log.Entry) ([]byte, error) {
	timestamp := time.Now().Format(time.RFC3339)
	hostname, _ := os.Hostname()
	return []byte(fmt.Sprintf("%s %s %s[%d]: %s %s\n", timestamp, hostname, tag, os.Getpid(), strings.ToUpper(entry.Level.String()), entry.Message)), nil
}

// tag represents the application name generating the log message. The tag
// string will appear in all log entires.
var tag string

func init() {
	tag = os.Args[0]
	log.SetFormatter(&ConfdFormatter{})
	logLevel := 4 // 0-panic, 2-error, 4-Info,  5-debug
	logPath := "./log/server.log"

	log.SetLevel(log.Level(logLevel))
	log.SetReportCaller(true)
	logger := &lumberjack.Logger{
		LocalTime:  true,
		Filename:   logPath,
		MaxSize:    20, // 一个文件最大为10M
		MaxBackups: 50, // 最多同时保存50份文件
		MaxAge:     60, // 一个文件最多同时存在60天
		Compress:   true,
	}
	writers := []io.Writer{
		logger,
		os.Stdout,
	}
	fileAndStdoutWriter := io.MultiWriter(writers...) //
	log.SetOutput(fileAndStdoutWriter)
}

// SetTag sets the tag.
func SetTag(t string) {
	tag = t
}

// SetLevel sets the log level. Valid levels are panic, fatal, error, warn, info and debug.
func SetLevel(level string) {
	lvl, err := log.ParseLevel(level)
	if err != nil {
		Fatal(fmt.Sprintf(`not a valid level: "%s"`, level))
	}
	log.SetLevel(lvl)
}

// Debug logs a message with severity DEBUG.
func Debug(format string, v ...interface{}) {
	log.Debug(fmt.Sprintf(format, v...))
}

// Error logs a message with severity ERROR.
func Error(format string, v ...interface{}) {
	log.Error(fmt.Sprintf(format, v...))
}

// Fatal logs a message with severity ERROR followed by a call to os.Exit().
func Fatal(format string, v ...interface{}) {
	log.Fatal(fmt.Sprintf(format, v...))
}

// Info logs a message with severity INFO.
func Info(format string, v ...interface{}) {
	log.Info(fmt.Sprintf(format, v...))
}

// Warning logs a message with severity WARNING.
func Warning(format string, v ...interface{}) {
	log.Warning(fmt.Sprintf(format, v...))
}
