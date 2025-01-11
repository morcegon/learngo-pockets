package pocketlog

import (
	"fmt"
	"io"
	"os"
)

// Logger is used to log information
type Logger struct {
	threshold    Level
	output       io.Writer
	messageLimit uint16
}

func New(threshold Level, opts ...Option) *Logger {
	lgr := &Logger{threshold: threshold, output: os.Stdout, messageLimit: 1000}

	for _, configFunc := range opts {
		configFunc(lgr)
	}

	return lgr
}

// Debugf formats and priunts a message if the log level is debug or higher
func (l *Logger) Debugf(format string, args ...any) {
	if l.threshold > LevelDebug {
		return
	}

	l.Logf(format, args...)
}

// Infof formats and prins a message if the log level is info or higher
func (l *Logger) Infof(format string, args ...any) {
	if l.threshold > LevelInfo {
		return
	}

	l.Logf(format, args...)
}

// Errorf formats and prins a message if the log level is error or higher
func (l *Logger) Errorf(format string, args ...any) {
	if l.threshold > LevelError {
		return
	}

	l.Logf(format, args...)
}

// Logf prints the message to the output
func (l *Logger) Logf(format string, args ...any) {
	if l.output == nil {
		l.output = os.Stdout
	}

	if l.messageLimit > 0 {
		format = l.limitMessageSize(format)
	}

	format = l.addLogLevel(format)

	_, _ = fmt.Fprintf(l.output, format+"\n", args...)
}

func (l *Logger) addLogLevel(format string) string {
	return "[" + logLevelName[l.threshold] + "] " + format
}

func (l *Logger) limitMessageSize(format string) string {
	for len(format) > int(l.messageLimit) {
		format = format[:l.messageLimit]
	}

	return format
}
