package main

import (
	"fmt"
)

// ANSI color codes
var ansiColors = map[string]string{
	"red":     "\033[31m",
	"green":   "\033[32m",
	"yellow":  "\033[33m",
	"blue":    "\033[34m",
	"magenta": "\033[35m",
	"cyan":    "\033[36m",
	"orange":  "\033[38;5;214m",
	"reset":   "\033[0m",
}

func PrintAsciiArtRecursive(sentences []string, textFile []string, color, substring string) {
	if len(sentences) == 0 {
		return
	}
	if sentences[0] != "" {
		printSentenceAsciiRecursive(sentences[0], textFile, 1, color, substring)
	} else {
		fmt.Println()
	}
	PrintAsciiArtRecursive(sentences[1:], textFile, color, substring)
}

func printSentenceAsciiRecursive(word string, textFile []string, h int, color, substring string) {
	if h > 8 {
		return
	}

	// Flag to track if any substring match was found
	foundSubstring := false

	i := 0
	for i < len(word) {
		// Check if the substring matches at the current position
		if substring != "" && i+len(substring) <= len(word) && word[i:i+len(substring)] == substring {
			// Color the substring characters
			for j := 0; j < len(substring); j++ {
				for lineIndex, line := range textFile {
					if lineIndex == (int(word[i+j])-32)*9+h {
						fmt.Print(ansiColors[color] + line + ansiColors["reset"])
					}
				}
			}
			foundSubstring = true // Set flag since we found the substring
			i += len(substring)   // Move index past the colored substring
		} else {
			// Print normally for non-matching characters
			for lineIndex, line := range textFile {
				if lineIndex == (int(word[i])-32)*9+h {
					// Only color the whole line if substring is empty (no substring mode)
					if substring == "" {
						fmt.Print(ansiColors[color] + line + ansiColors["reset"])
					} else {
						fmt.Print(line)
					}
				}
			}
			i++
		}
	}

	// Color the whole string if no substring was matched and a substring was specified
	if substring != "" && !foundSubstring {
		// Print the whole string in color if substring wasn't found
		for i := 0; i < len(word); i++ {
			for lineIndex, line := range textFile {
				if lineIndex == (int(word[i])-32)*9+h {
					fmt.Print(ansiColors[color] + line + ansiColors["reset"])
				}
			}
		}
	}

	fmt.Println()
	// Recursive call for the next line height
	printSentenceAsciiRecursive(word, textFile, h+1, color, substring)
}
