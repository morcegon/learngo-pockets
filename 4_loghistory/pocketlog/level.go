package pocketlog

// Level represents an available logging level
type Level byte

const (
	// LevelDebug represents the lowest level of log, mostly used for debugging purposes
	LevelDebug Level = iota
	// LevelInfo represents a logging level that contains information deemed valuable
	LevelInfo Level = iota
	// LevelError represents the highest logging level, only to be used to trace errors
	LevelError Level = iota
)

var logLevelName = map[Level]string{
	LevelDebug: "Debug",
	LevelInfo:  "Info",
	LevelError: "Error",
}
