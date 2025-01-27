package gordle

import "testing"

func inCorpus(corpus []string, word string) bool {
	for _, corpusWord := range corpus {
		if corpusWord == word {
			return true
		}
	}

	return false
}

func TestPickWorkd(t *testing.T) {
	corpus := []string{"HELOO", "SALUT", "XAIPE", "COISA"}
	word := pickWord(corpus)

	if !inCorpus(corpus, word) {
		t.Errorf("expected a word in corpus, got %q", word)
	}
}
