package gordle

import (
	"errors"
	"slices"
	"strings"
	"testing"
)

func TestGameAsk(t *testing.T) {
	tt := map[string]struct {
		input string
		want  []rune
	}{
		"5 characters in english": {
			input: "HELLO",
			want:  []rune("HELLO"),
		},
		"5 characters in arabic": {
			input: "مرحبا",
			want:  []rune("مرحبا"),
		},
		"5 characters in japanese": {
			input: "こんにちは",
			want:  []rune("こんにちは"),
		},
		"3 characters in japanese": {
			input: "こんに\nこんにちは",
			want:  []rune("こんにちは"),
		},
		"to uppercase": {
			input: "lower",
			want:  []rune("LOWER"),
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			g := New(strings.NewReader(tc.input), string(tc.want), 0)

			got := g.ask()
			if !slices.Equal(got, tc.want) {
				t.Errorf("got = %v, want %v", string(got), string(tc.want))
			}
		})
	}
}

func TestValidateGuess(t *testing.T) {
	tt := map[string]struct {
		word     []rune
		guess    []rune
		expected error
	}{
		"nominal": {
			word:     []rune("GUESS"),
			guess:    []rune("GUESS"),
			expected: nil,
		},
		"too long": {
			word:     []rune("POCKE"),
			guess:    []rune("POCKET"),
			expected: errInvalidWordLenght,
		},
		"too short": {
			word:     []rune("TINYA"),
			guess:    []rune("TINY"),
			expected: errInvalidWordLenght,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			g := New(nil, string(tc.word), 0)

			err := g.validateGuess(tc.guess)
			if !errors.Is(err, tc.expected) {
				t.Errorf("%c, expxected %q, got %q", tc.word, tc.expected, err)
			}
		})
	}
}

func TestComputeFeedback(t *testing.T) {
	tt := map[string]struct {
		guess            string
		solution         string
		expectedFeedback feedback
	}{
		"nominal": {
			guess:            "ASDFG",
			solution:         "ASDFG",
			expectedFeedback: feedback{2, 2, 2, 2, 2},
		},
		"double character": {
			guess:            "HELLO",
			solution:         "HELLO",
			expectedFeedback: feedback{2, 2, 2, 2, 2},
		},
		"double with wrong answer": {
			guess:            "HELOO",
			solution:         "HELLO",
			expectedFeedback: feedback{2, 2, 2, 0, 2},
		},
		"two identical, but not in the right position": {
			guess:            "HELOL",
			solution:         "HELLO",
			expectedFeedback: feedback{2, 2, 2, 1, 1},
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			fb := computeFeedback([]rune(tc.guess), []rune(tc.solution))
			if !tc.expectedFeedback.Equal(fb) {
				t.Errorf(
					"guess: %q, got the wrong feedback, expxected %v, got %v",
					tc.guess,
					tc.expectedFeedback,
					fb,
				)
			}
		})
	}
}
