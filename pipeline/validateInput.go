package pipeline

import (
	"errors"
	"fmt"
)

// ValidateInput checks that the input string meets the following requirements:
// 1. The input is not empty.
// 2. The input does not exceed 10,000 runes (characters) in length.
// 3. The input does not contain control characters other than tab (\t), newline (\n), and carriage return (\r).
// If any condition is violated, it returns a non-nil error describing the problem.
func ValidateInput(input string) error {
	// Count the total number of runes (Unicode characters) in the input string.
	// This accounts for multi-byte UTF-8 sequences as single runes.
	runeCount := 0
	for range input {
		runeCount++
	}

	// Check if the input is empty (zero runes). If so, return an error.
	if runeCount == 0 {
		return errors.New("input is empty")
	}

	// Check if the input exceeds the maximum allowed length of 10,000 runes.
	if runeCount > 10000 {
		return errors.New("input too long")
	}

	// Iterate through each rune in the input string to validate individual characters.
	for _, r := range input {
		// If the rune is a control character (codepoint below 32):
		if r < 32 {
			// Allow three specific control characters: tab (9), newline (10), and carriage return (13).
			if r != '\t' && r != '\n' && r != '\r' {
				// If the control character is not one of the allowed three, reject it.
				return fmt.Errorf("invalid control character: 0x%x", r)
			}
		}
	}

	// If all validations passed, return nil to indicate the input is valid.
	return nil
}
