package main

import (
	"runtime"
	"strings"

	"github.com/fatih/color"
)

// ColorString - Returns colored string
func (l *logger) ColorString(s, col string) string {

	// If Windows, disable colorS
	if runtime.GOOS == "windows" || *l.Colors {
		return s
	}

	var result string
	switch strings.ToLower(col) {
	case "green":
		result = color.New(color.FgGreen).Add().SprintFunc()(s)
	case "yellow":
		result = color.New(color.FgYellow).Add().SprintFunc()(s)
	case "red":
		result = color.New(color.FgRed).Add().SprintFunc()(s)
	case "magenta":
		result = color.New(color.FgMagenta).Add().SprintFunc()(s)
	case "cyan":
		result = color.New(color.FgCyan).Add().SprintFunc()(s)
	default:
		// Unidentified, just returns the same string
		return s
	}

	return result
}
