package pipeline

import "strings"

func RenderLines(tokens []string, banner map[string][]string) []string {
	// Initialize the output slice to collect all rendered lines (8 lines per character).
	var out []string
	// Initialize a buffer of 8 strings to hold the current row being rendered.
	current := make([]string, 8)
	// Track whether we've rendered any characters since the last newline.
	hasContent := false
	// Track if the last token was a newline to handle consecutive newlines (blank lines).
	lastWasNewline := false

	// flush writes the current 8-line buffer to the output and resets it for the next character.
	flush := func() {
		// Only append if we have actual content to flush (skip consecutive newlines).
		if hasContent {
			// Append the current buffer of 8 lines to the output slice.
			out = append(out, current...)
			// Reset the buffer to prepare for the next set of 8 lines.
			current = make([]string, 8)
			// Mark that we've flushed the content.
			hasContent = false
		}
	}

	for _, tok := range tokens {
		if tok == "\n" {
			flush()
			// If the last token was also a newline, add a blank line (consecutive newlines = blank line).
			if lastWasNewline {
				out = append(out, " ")
			}
			lastWasNewline = true
			continue
		}
		// Any non-newline character means we're no longer in consecutive newline territory.
		lastWasNewline = false

		// Look up the glyph (ASCII art representation) for the current token in the banner map.
		glyph, ok := banner[tok]
		if !ok {
			// If the token is not found in the banner, use 4 spaces as a placeholder.
			pad := strings.Repeat(" ", 4)
			// Add the padding to each of the 8 rows to maintain alignment.
			for i := 0; i < 8; i++ {
				current[i] += pad
			}
			// Mark that we have content.
			hasContent = true
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
		// Mark that we have content.
		hasContent = true
	}

	// After processing all tokens, flush the final buffer to the output.
	flush()
	// Return the complete list of 8-line rows representing the rendered ASCII art.
	return out
}
