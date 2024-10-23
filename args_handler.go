package main

import (
	"fmt"
	"regexp"
)

// Argument handling function
func handleArguments(args []string) (string, string, string, error) {
	// Check the number of arguments
	if len(args) < 2 || len(args) > 5 {
		return "", "", "", fmt.Errorf("usage: go run . [OPTION] [STRING]\nEX: go run . --color=<color> <substring to be colored> 'something'")
	}

	var banner string
	if len(args) == 5 {
		banner = args[3]
		switch banner {
		case "standard":
			banner = "banner/standard.txt"
		case "shadow":
			banner = "banner/shadow.txt"
		case "thinkertoy":
			banner = "banner/thinkertoy.txt"
		default:
			return "", "", "", fmt.Errorf("use a valid banner template")
		}
	} else {
		banner = "banner/standard.txt"
	}

	// Check for color flag
	re := regexp.MustCompile(`--color=(\w+)`)
	flag := args[1]
	matches := re.FindStringSubmatch(flag)
	if matches == nil {
		return "", "", "", fmt.Errorf("must input correct color format (--color=<color>)")
	}

	color := matches[1]
	// Check if the color is valid
	if _, ok := colorCodes[color]; !ok {
		return "", "", "", fmt.Errorf("invalid color. Available colors: red, green, yellow, blue, magenta, cyan")
	}

	// Get the substring if provided
	substring := ""
	if len(args) == 4 {
		substring = args[2]
	}

	return banner, color, substring, nil
}
