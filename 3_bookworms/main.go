package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var filePath string
	flag.StringVar(&filePath, "filePath", "", "The required bookworms file path")
	flag.Parse()

	bookworms, err := loadBookWorkms(filePath)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "failed to load bookworms: %s\n", err)
		os.Exit(1)
	}

	books := findCommonBooks(bookworms)

	fmt.Println("Here are the books in common:")
	displayBooks(books)
}
