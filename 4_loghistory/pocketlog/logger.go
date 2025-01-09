package pocketlog

import (
	"fmt"
	"io"
	"os"
)

// Logger is used to log information
type Logger struct {
	threshold Level
	output    io.Writer
}

func New(threshold Level, output io.Writer) *Logger {
	return &Logger{
		threshold: threshold,
		output:    output,
	}
}

// Debugf formats and priunts a message if the log level is debug or higher
func (l *Logger) Debugf(format string, args ...any) {
	if l.threshold > LevelDebug {
		return
	}

	l.logf(format, args...)
}

// Info prints a message if the log level is info or higher
func (l *Logger) Info(args ...any) {
	if l.threshold > LevenInfo {
		return
	}

	_, _ = fmt.Print(args...)
}

// Infof formats and prins a message if the log level is info or higher
func (l *Logger) Infof(format string, args ...any) {
	if l.threshold > LevenInfo {
		return
	}

	_, _ = fmt.Printf(format, args...)
}

// Error prins a message if the log level is error or higher
func (l *Logger) Error(args ...any) {
	if l.threshold > LevelError {
		return
	}

	_, _ = fmt.Print(args...)
}

// Errorf formats and prins a message if the log level is error or higher
func (l *Logger) Errorf(format string, args ...any) {
	if l.threshold > LevelError {
		return
	}

	_, _ = fmt.Printf(format, args...)
}

// logf prints the message to the output
func (l *Logger) logf(format string, args ...any) {
	if l.output == nil {
		l.output = os.Stdout
	}

	_, _ = fmt.Fprintf(l.output, format+"\n", args...)
}
