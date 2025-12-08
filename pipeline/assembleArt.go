package pipeline

import (
	"strings"
)

// AssembleArt joins a slice of line strings into a single output string,
// using newline characters ("\n") as separators between lines.
// No trailing newline is appended at the end.
// If the input slice is empty, it returns an empty string.
func AssembleArt(lines []string) string {
	// If the input slice has no lines, return an empty string immediately.
	if len(lines) == 0 {
		return ""
	}

	// Join all lines together using the newline character as the separator.
	// This concatenates the lines with "\n" between each pair.
	return strings.Join(lines, "\n")
}
