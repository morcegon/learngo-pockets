package pocketlog

import "io"

type Option func(l *Logger)

// Returns a configuration function that sets the ouput of a log
func WithOutput(output io.Writer) Option {
	return func(lgr *Logger) {
		lgr.output = output
	}
}

// Trim messages at an specified limit. 1000 is the default
func LimitMessages(limit uint16) Option {
	return func(lgr *Logger) {
		lgr.messageLimit = limit
	}
}
