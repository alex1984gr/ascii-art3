package pipeline

import (
	"errors"
	"io"
	"strings"
)

// WriteOutput writes a slice of strings (ASCII art lines) to the given io.Writer.
// If a line is the literal "\n", it writes an actual newline instead.
func WriteOutput(lines []string, w io.Writer) error {
	// Check if the writer is nil; return an error if it is
	if w == nil {
		return errors.New("writer is nil")
	}

	// If there are no lines to write, do nothing and return nil
	if len(lines) == 0 {
		return nil
	}

	// Trim leading spaces from the first line to avoid unwanted indentation
	lines[0] = strings.TrimLeft(lines[0], " ")

	// Iterate over each line in the slice
	for _, line := range lines {
		// If the line is the literal string "\n", write a real newline to the writer
		if line == `\n` {
			_, err := io.WriteString(w, "\n")
			if err != nil {
				return err // Return the error if writing fails
			}
			continue // Move to the next line
		}

		// Otherwise, write the line followed by a newline
		_, err := io.WriteString(w, line+"\n")
		if err != nil {
			return err // Return the error if writing fails
		}
	}

	// Successfully wrote all lines
	return nil
}
