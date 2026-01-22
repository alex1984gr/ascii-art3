package pipeline

import "strings"

func RenderLines(tokens []string, banner map[string][]string) []string {
	// Initialize the output slice to collect all rendered lines (8 lines per character).
	var out []string
	// Initialize a buffer of 8 strings to hold the current row being rendered.
	current := make([]string, 8)

	// flush writes the current 8-line buffer to the output and resets it for the next character.
	flush := func() {
		// Append the current buffer of 8 lines to the output slice.
		// We always append even if lines contain only spaces, as spaces are part of ASCII art.
		out = append(out, current...)
		// Reset the buffer to prepare for the next set of 8 lines.
		current = make([]string, 8)
	}

	for _, tok := range tokens {
		if tok == "\n" {
			flush()
			continue
		}

		// Look up the glyph (ASCII art representation) for the current token in the banner map.
		glyph, ok := banner[tok]
		if !ok {
			// If the token is not found in the banner, use 4 spaces as a placeholder.
			pad := strings.Repeat(" ", 4)
			// Add the padding to each of the 8 rows to maintain alignment.
			for i := 0; i < 8; i++ {
				current[i] += pad
			}
			// Move to the next token without processing further.
			continue
		}

		// Ensure the glyph has exactly 8 rows to maintain consistent banner format.
		if len(glyph) < 8 {
			// Create a new slice with capacity for exactly 8 rows.
			tmp := make([]string, 8)
			// Copy the existing glyph rows to the new slice (remaining positions stay empty).
			copy(tmp, glyph)
			// Use the padded glyph.
			glyph = tmp
		}

		// Append each row of the glyph to the corresponding row in the current buffer.
		for i := 0; i < 8; i++ {
			// Concatenate the glyph row to the current row being built.
			current[i] += glyph[i]
		}
	}

	// After processing all tokens, flush the final buffer to the output.
	flush()
	// Return the complete list of 8-line rows representing the rendered ASCII art.
	return out
}
