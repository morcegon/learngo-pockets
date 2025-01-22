package gordle

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"slices"
	"strings"
)

type Game struct {
	reader      *bufio.Reader
	solution    []rune
	maxAttempts int
}

func New(playerInput io.Reader, solution string, maxAttempts int) *Game {
	g := &Game{
		reader:      bufio.NewReader(playerInput),
		solution:    splitToUppercaseCharacters(solution),
		maxAttempts: maxAttempts,
	}

	return g
}

func (g *Game) Play() {
	fmt.Println("Welcome to Gordle")

	solution := string(g.solution)

	for currentAttempt := 1; currentAttempt <= g.maxAttempts; currentAttempt++ {
		guess := g.ask()

		fb := computeFeedback(guess, g.solution)

		fmt.Println(fb.String())

		if slices.Equal(guess, g.solution) {
			fmt.Printf(
				"🎉 You won! You found it in %d guess(es)! The word was: %s.\n",
				currentAttempt,
				solution,
			)
			return
		}
	}

	fmt.Printf("😬 You've lost! The solution was: %s.\n", solution)
}

// ask reads input until a valid suggestion is made (and returned)
func (g *Game) ask() []rune {
	fmt.Printf("Enter a %d-character guess:\n", len(g.solution))
	for {
		playerInput, _, err := g.reader.ReadLine()
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Gordle failed to read your guess: %s\n", err.Error())
		}

		guess := splitToUppercaseCharacters(string(playerInput))

		err = g.validateGuess(guess)
		if err != nil {
			_, _ = fmt.Fprintf(
				os.Stderr,
				"Your attempt is invalid with Gordle's solution: %s\n",
				err.Error(),
			)
		} else {
			return guess
		}

	}
}

var errInvalidWordLenght = fmt.Errorf(
	"invalid guess, word doesn't have the same number of characters as the solution",
)

// ensures the guess is valid enough
func (g *Game) validateGuess(guess []rune) error {
	guessLenght := len(guess)
	solutionLenght := len(g.solution)

	if guessLenght != solutionLenght {
		return fmt.Errorf(
			"expected %d got %d, %w",
			solutionLenght,
			guessLenght,
			errInvalidWordLenght,
		)
	}

	return nil
}

func splitToUppercaseCharacters(input string) []rune {
	return []rune(strings.ToUpper(input))
}

func computeFeedback(guess, solution []rune) feedback {
	result := make(feedback, len(guess))
	used := make([]bool, len(solution))

	if len(guess) != len(solution) {
		_, _ = fmt.Fprintf(
			os.Stderr,
			"Internal error! Guess and solution have different lengths: %d vs %d",
			len(guess),
			len(solution),
		)
		return result
	}

	// check the correct letters
	for posIngGuess, character := range guess {
		if character == solution[posIngGuess] {
			result[posIngGuess] = correctPostion
			used[posIngGuess] = true
		}
	}

	// look for letters in wrong correctPostiion
	for posIngGuess, character := range guess {
		if result[posIngGuess] != absentCharacter {
			continue
		}

		for posInSolution, target := range solution {
			if used[posInSolution] {
				// The letter of the solution is already assigned to a letter of the guess
				// Skip to the next letter of the solution
				continue
			}

			if character == target {
				result[posIngGuess] = wrongPosition
				used[posInSolution] = true
				// Skip to the next letter of the guess
				break
			}
		}
	}

	return result
}
