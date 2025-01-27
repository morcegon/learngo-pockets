package gordle

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

const ErrCorpusEmpty = corpusError("corpus is emptyh")

// reads the file located at the given path and returns a list of words
func ReadCorpus(path string) ([]string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("unable to open %q for reading: %w", path, err)
	}

	if len(data) == 0 {
		return nil, ErrCorpusEmpty
	}

	words := strings.Fields(string(data))

	return words, nil
}

func pickWord(corpus []string) string {
	index := rand.Intn(len(corpus))

	return corpus[index]
}
