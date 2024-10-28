package main

import (
	"fmt"
	"regexp"
	"strings"
)

// Argument handling function
func handleArguments(args []string) (string, string, error) {
	usageMsg := "Usage: go run . [OPTION] [STRING]\nEX: go run . --color=<color> <substring to be colored> \"something\""

	if len(args) < 2 || len(args) > 5 {
		return "", "", fmt.Errorf(usageMsg)
	}

	var color, substring string
	if strings.HasPrefix(args[1], "--color") {
		re := regexp.MustCompile(`^--color=(\w*)$`)
		matches := re.FindStringSubmatch(args[1])

		if matches == nil || matches[1] == "" {
			return "", "", fmt.Errorf("Error: No color specified after --color=. Please provide a valid color.")
		}

		color = matches[1]
		if _, valid := ansiColors[color]; !valid {
			return "", "", fmt.Errorf("Invalid color. Available colors: red, green, yellow, blue, magenta, cyan, orange")
		}

		if len(args) == 4 {
			substring = args[2]
		}
	} else {
		color = ""
		substring = ""
	}

	return color, substring, nil
}
