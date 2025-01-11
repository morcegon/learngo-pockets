package pocketlog_test

import (
	"os"
	"testing"

	"learngo-pockets/logger/pocketlog"
)

const (
	debugMessage = "Why wirte I still all one, ever the same,"
	infoMessage  = "And keep invention in a noted weed,"
	errorMessage = "That every word doth almost twell my name,"
)

func ExampleLogger_debugf() {
	debugLogger := pocketlog.New(pocketlog.LevelDebug, pocketlog.WithOutput(os.Stdout))
	debugLogger.Debugf("Hello, %s", "world")
	// Output: [Debug] Hello, world
}

func TestLogger_DebugInfoErrorf(t *testing.T) {
	type testCase struct {
		level    pocketlog.Level
		expected string
	}

	tt := map[string]testCase{
		"debug": {
			level:    pocketlog.LevelDebug,
			expected: "[Debug] " + debugMessage + "\n" + "[Debug] " + infoMessage + "\n" + "[Debug] " + errorMessage + "\n",
		},
		"info": {
			level:    pocketlog.LevelInfo,
			expected: "[Info] " + infoMessage + "\n" + "[Info] " + errorMessage + "\n",
		},
		"error": {
			level:    pocketlog.LevelError,
			expected: "[Error] " + errorMessage + "\n",
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			tw := &testWriter{}

			testedLogger := pocketlog.New(tc.level, pocketlog.WithOutput(tw))

			testedLogger.Debugf(debugMessage)
			testedLogger.Infof(infoMessage)
			testedLogger.Errorf(errorMessage)

			if tw.contents != tc.expected {
				t.Errorf("invalid contents, expected %q, got %q", tc.expected, tw.contents)
			}
		})
	}
}

func TestLogger_LimitMessage(t *testing.T) {
	tt := map[string]struct {
		limit    uint16
		input    string
		expected string
	}{
		"limit": {
			limit:    2,
			input:    "Hello world",
			expected: "[Debug] He\n",
		},
		"default limit": {
			input:    "Hello world",
			expected: "[Debug] Hello world\n",
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			tw := &testWriter{}

			logger := pocketlog.New(
				pocketlog.LevelDebug,
				pocketlog.WithOutput(tw),
				pocketlog.LimitMessages(tc.limit),
			)

			logger.Debugf(tc.input)

			if tw.contents != tc.expected {
				t.Fatalf("invalid content, expected %q, got %q", tc.expected, tw.contents)
			}
		})
	}
}

type testWriter struct {
	contents string
}

func (tw *testWriter) Write(p []byte) (n int, err error) {
	tw.contents = tw.contents + string(p)
	return len(p), nil
}
