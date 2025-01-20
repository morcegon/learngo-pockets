package main

import (
	"os"

	"learngo-pockets/gordle/gordle"
)

const maxAttemtps = 6

func main() {
	solution := "hello"
	g := gordle.New(os.Stdin, solution, maxAttemtps)
	g.Play()
}
