package main

import (
	"fmt"
	"os"

	"learngo-pockets/gordle/gordle"
)

const maxAttemtps = 6

func main() {
	corpus, err := gordle.ReadCorpus("./corpus/english.txt")
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "unable to read corpus: %s", err)
		return
	}

	g, err := gordle.New(os.Stdin, corpus, maxAttemtps)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "unable to start the game: %s", err)
		return
	}

	g.Play()
}
