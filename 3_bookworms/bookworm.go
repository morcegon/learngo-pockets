package main

import (
	"encoding/json"
	"os"
)

type Bookworm struct {
	Name  string `json:"name"`
	Books []Book `json:"books"`
}

type Book struct {
	Author string `json:"author"`
	Title  string `json:"title"`
}

func loadBookWorkms(filePath string) ([]Bookworm, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	var bookworks []Bookworm

	// Decode the file and store the content in the value bookworms
	err = json.NewDecoder(f).Decode(&bookworks)
	if err != nil {
		return nil, err
	}

	return bookworks, nil
}

func findCommonBooks(bookworms []Bookworm) []Bookworm {
	return nil
}

func booksCount(bookworms []Bookworm) map[Book]uint {
	count := make(map[Book]uint)

	for _, bookworm := range bookworms {
		for _, book := range bookworm.Books {
			count[book]++
		}
	}

	return count
}
