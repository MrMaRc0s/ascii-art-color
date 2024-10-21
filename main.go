package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	// Ensure correct argument usage
	if len(os.Args) < 1 || len(os.Args) > 5 {
		fmt.Println("Usage: go run . [OPTION] [STRING] \n EX: go run . --color=<color> <substring to be colored> 'something'")
		os.Exit(1)
	}

	var banner string
	if len(os.Args) == 5 {
		banner = os.Args[3]
		switch banner {
		case "standard":
			banner = "banner/standard.txt"
		case "shadow":
			banner = "banner/shadow.txt"
		case "thinkertoy":
			banner = "banner/thinkertoy.txt"
		default:
			fmt.Println("Use a valid banner template")
			os.Exit(2)
		}
	}

	if len(os.Args) == 2 || len(os.Args) == 3 {
		banner = "banner/standard.txt"
	}

	re := regexp.MustCompile(`--color=`)

	flag := os.Args[1]

	if !re.MatchString(flag) {
		fmt.Println("Must input correct color or color format")
		os.Exit(3)
	}

	// Read the file
	bannerFile, err := os.ReadFile(banner)
	if err != nil {
		fmt.Println("There was a problem reading the banner file template")
		os.Exit(3)
	}

	// Split the lines of the file into a slice
	line := strings.Replace(string(bannerFile), "\r\n", "\n", -1)
	lines := strings.Split(line, "\n")

	// Split the input string into lines, based on literal newline characters
	sepText := strings.Split(os.Args[1], "\\n") // Handling literal "\n" from the input

	// Print the ASCII Art based on the input
	PrintAsciiArtRecursive(sepText, lines)
}
