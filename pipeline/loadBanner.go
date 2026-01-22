package pipeline

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// LoadBanner loads a banner font file by name from the banners directory.
// The name parameter should be the banner name without the .txt extension (e.g., "standard").
// It returns a map where keys are character strings and values are slices of 8 rows (lines of ASCII art).
// If the file does not exist or cannot be parsed, it returns a non-nil error.
func LoadBanner(name string) (map[string][]string, error) {
	// Check that the banner name is not empty; return error if it is.
	if name == "" {
		return nil, errors.New("empty banner name")
	}

	// Build the file path "banners/<name>.txt" using the platform-appropriate path separator.
	path := filepath.Join("banners", name+".txt")

	// Try to open the file at this path.
	f, err := os.Open(path)
	if err != nil {
		// If opening fails, try "../banners/<name>.txt" (useful when running tests from `tests/` folder).
		path = filepath.Join("..", "banners", name+".txt")
		var err2 error
		f, err2 = os.Open(path)
		if err2 != nil {
			// If both attempts fail, return the original error.
			return nil, fmt.Errorf("open banner file: %w", err)
		}
	}
	// Ensure the file is closed when the function exits, even if an error occurs.
	defer f.Close()

	// Parse the opened file using the common reader parser.
	return LoadBannerFromReader(f)
}

// LoadBannerFromReader parses a banner font from an io.Reader in compact format.
//
// The format is:
//
//	CHAR:<character or escape sequence>
//	<row 1>
//	<row 2>
//	...
//	<row 8>
//
// Escape sequences like "\n", "\t", "\r" are unescaped to their actual characters.
// Returns a map of character keys to 8-row ASCII art glyphs.
func LoadBannerFromReader(r io.Reader) (map[string][]string, error) {
	// Read all lines from the reader using a scanner to process line-by-line.
	scanner := bufio.NewScanner(r)
	var lines []string
	// Iterate through each line in the reader and collect them.
	for scanner.Scan() {
		// Append the current line to the lines slice.
		lines = append(lines, scanner.Text())
	}
	// Check if any error occurred during scanning.
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	// Detect if the file uses the "CHAR:" header format by searching for the prefix.
	usesCharHeader := false
	for _, l := range lines {
		// Check if any line starts with "CHAR:".
		if strings.HasPrefix(l, "CHAR:") {
			usesCharHeader = true
			break
		}
	}

	// Prepare the map to hold character => 8-row ASCII art.
	banner := make(map[string][]string)

	if usesCharHeader {
		// Parse CHAR:<literal> format (explicit character headers).
		var currentKey string // Stores the current character being parsed.
		var rows []string     // Stores the 8 rows for that character.
		// Define a nested function to convert escape sequences to actual characters.
		unescape := func(raw string) string {
			// Convert literal escape sequences to actual characters.
			switch raw {
			case "\\n":
				return "\n"
			case "\\t":
				return "\t"
			case "\\r":
				return "\r"
			case "\\\\":
				return "\\"
			default:
				return raw
			}
		}

		// Iterate over each line in the file.
		for _, line := range lines {
			if strings.HasPrefix(line, "CHAR:") {
				// If we already have a character being parsed, finalize it before starting a new one.
				if currentKey != "" {
					// Pad rows to exactly 8 if there are fewer.
					for len(rows) < 8 {
						rows = append(rows, "")
					}
					// Copy rows to banner map (create a new slice to avoid reference issues).
					banner[currentKey] = append([]string(nil), rows...)
				}
				// Extract the character key after "CHAR:" and unescape any escape sequences.
				rawKey := strings.TrimPrefix(line, "CHAR:")
				currentKey = unescape(rawKey)
				// Reset rows for the new character.
				rows = nil
				continue
			}

			// Skip lines if we haven't encountered a CHAR: header yet.
			if currentKey == "" {
				continue
			}
			// Append this line to the current character's rows without trimming trailing spaces.
			// Trailing spaces are essential to the character's width in the ASCII art design.
			rows = append(rows, line)
		}

		// After finishing all lines, store the last character being parsed.
		if currentKey != "" {
			// Pad rows to exactly 8 if needed.
			for len(rows) < 8 {
				rows = append(rows, "")
			}
			// Add the final character to the banner map.
			banner[currentKey] = append([]string(nil), rows...)
		}

		// Return the parsed banner map.
		return banner, nil
	}

	// If no CHAR: headers, parse blank-line-separated blocks format (legacy format).
	var block []string // Temporary storage for the current glyph being built.
	char := rune(31)   // Start before space (ASCII 32) to begin with space on first increment.
	// Define a nested function to emit a completed block to the banner map.
	emitBlock := func() {
		// Skip empty blocks (no content).
		if len(block) == 0 {
			return
		}
		// Increment to the next ASCII character.
		char++
		// Ensure the block has exactly 8 rows for compatibility with the banner format.
		for len(block) < 8 {
			block = append(block, "")
		}
		// Add the block to the banner map using the character as a string key.
		banner[string(char)] = append([]string(nil), block...)
		// Reset block for the next character.
		block = nil
	}

	// Iterate over all lines in the file.
	for _, line := range lines {
		// Check if the line is blank (whitespace only).
		if strings.TrimSpace(line) == "" {
			// Blank line signals the end of the current glyph block.
			emitBlock()
			continue
		}
		// Add the non-blank line to the current block.
		block = append(block, line)
	}
	// Emit any trailing block that was being built at the end of the file.
	emitBlock()

	// Return the completed banner map with all characters parsed.
	return banner, nil
}
