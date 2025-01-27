package gordle_test

import (
	"testing"

	"learngo-pockets/gordle/gordle"
)

func TestReadCorpus(t *testing.T) {
	tt := map[string]struct {
		file   string
		length int
		err    error
	}{
		"English corpus": {
			file:   "../corpus/english.txt",
			length: 97,
			err:    nil,
		},
		"Empty corpus": {
			file:   "../corpus/empty.txt",
			length: 0,
			err:    gordle.ErrCorpusEmpty,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			words, err := gordle.ReadCorpus(tc.file)
			if tc.err != err {
				t.Errorf("expected err %v, got %v", tc.err, err)
			}

			if tc.length != len(words) {
				t.Errorf("expected %d, got %d", tc.length, len(words))
			}
		})
	}
}
