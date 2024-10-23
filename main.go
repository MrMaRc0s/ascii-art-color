package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// Argument handling
	banner, color, substring, err := handleArguments(os.Args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Read the banner file
	bannerFile, err := os.ReadFile(banner)
	if err != nil {
		fmt.Println("There was a problem reading the banner file template")
		os.Exit(3)
	}

	// Split the lines of the file into a slice
	line := strings.Replace(string(bannerFile), "\r\n", "\n", -1)
	lines := strings.Split(line, "\n")

	// Split the input string into lines, based on literal newline characters
	sepText := strings.Split(os.Args[len(os.Args)-1], "\\n")

	// Print ASCII art with color
	PrintAsciiArtRecursive(sepText, lines, color, substring)
}
