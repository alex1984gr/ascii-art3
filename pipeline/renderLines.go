package pipeline

import (
	"strings"
)

// RenderLines takes a slice of character tokens and a banner map, and produces
// a slice of 8 output lines representing the visual rendering of those tokens.
// Each token contributes its 8-row glyph from the banner map, concatenated horizontally.
// A newline token ("\n") causes the current line block to be appended to the output
// and a new block to be started (vertical stacking).
// If a token is not found in the banner, it is replaced with 3 spaces (default padding).
func RenderLines(tokens []string, banner map[string][]string) []string {
	// Initialize the output slice to collect all rendered lines.
	var out []string

	// Initialize a working array of 8 strings representing the current horizontal block.
	current := make([]string, 8)

	// Process each token in order.
	for _, tok := range tokens {
		// If the token is a newline, flush the current block and start a new one.
		if tok == "\n" {
			// Append the current 8 rows to the output.
			out = append(out, current...)

			// Reset the current block to a fresh 8-row array of empty strings.
			current = make([]string, 8)
			continue
		}

		// Look up the token in the banner map to get its glyph (8 rows).
		glyph, ok := banner[tok]

		// If the token is not found in the banner, use a fallback width of 4 spaces.
		if !ok {
			// Create a 4-character wide padding (all spaces).
			pad := strings.Repeat(" ", 4)

			// Append the padding to each of the 8 rows.
			for i := 0; i < 8; i++ {
				current[i] += pad
			}
			continue
		}

		// Ensure the glyph has exactly 8 rows (pad with empty strings if needed).
		for len(glyph) < 8 {
			glyph = append(glyph, "")
		}

		// Concatenate the glyph rows to the current rows (build horizontal).
		for i := 0; i < 8; i++ {
			current[i] += glyph[i]
		}
	}

	// After processing all tokens, append the final current block to the output.
	out = append(out, current...)

	// Return the complete set of rendered output lines.
	return out
}
