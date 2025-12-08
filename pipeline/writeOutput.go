package pipeline

import (
	"errors"
	"io"
	"strings"
)

// WriteOutput writes the provided lines to the given writer.
// Each line is written followed by a newline character.
// The first line has any leading whitespace trimmed.
// If there are no lines, nothing is written but no error is returned.
// If the writer is nil, it returns a non-nil error.
func WriteOutput(lines []string, w io.Writer) error {
	// Check if the writer is nil; if so, return an error immediately.
	if w == nil {
		return errors.New("writer is nil")
	}

	// If the input slice has no lines, return nil (nothing to write, no error).
	if len(lines) == 0 {
		return nil
	}

	// Make a copy of the lines so we don't modify the original slice.
	outputLines := make([]string, len(lines))
	copy(outputLines, lines)

	// Trim leading whitespace from the first line.
	outputLines[0] = strings.TrimLeft(outputLines[0], " ")

	// Join all lines using the newline character as a separator, then append one final newline.
	out := strings.Join(outputLines, "\n") + "\n"

	// Write the assembled string to the provided writer.
	// If an error occurs during writing, return it to the caller.
	_, err := io.WriteString(w, out)
	return err
}
