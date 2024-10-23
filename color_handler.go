package main

import "strings"

// ANSI color codes
var colorCodes = map[string]string{
	"red":     "\033[31m",
	"green":   "\033[32m",
	"yellow":  "\033[33m",
	"blue":    "\033[34m",
	"magenta": "\033[35m",
	"cyan":    "\033[36m",
	"reset":   "\033[0m",
}

// New coloring function
func ApplyColor(line, color, substring string) string {
	if substring == "" || strings.Contains(line, substring) {
		return colorCodes[color] + line + colorCodes["reset"]
	}
	return line
}
