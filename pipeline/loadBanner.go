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
	// Check that the banner name is not empty; return error if it is
	if name == "" {
		return nil, errors.New("empty banner name")
	}

	// Build the file path "banners/<name>.txt"
	path := filepath.Join("banners", name+".txt")

	// Try to open the file at this path
	f, err := os.Open(path)
	if err != nil {
		// If opening fails, try "../banners/<name>.txt" (useful when running tests from `tests/` folder)
		path = filepath.Join("..", "banners", name+".txt")
		var err2 error
		f, err2 = os.Open(path)
		if err2 != nil {
			// If both attempts fail, return the original error
			return nil, fmt.Errorf("open banner file: %w", err)
		}
	}
	// Ensure the file is closed when the function exits
	defer f.Close()

	// Parse the opened file using the common reader parser
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
	// Read all lines from the reader using a scanner
	scanner := bufio.NewScanner(r)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text()) // collect each line into a slice
	}
	// Return any scanning error
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	// Detect if the file uses the "CHAR:" header format
	usesCharHeader := false
	for _, l := range lines {
		if strings.HasPrefix(l, "CHAR:") {
			usesCharHeader = true
			break
		}
	}

	// Prepare the map to hold character => 8-row ASCII art
	banner := make(map[string][]string)

	if usesCharHeader {
		// Parse CHAR:<literal> format
		var currentKey string // stores the current character being parsed
		var rows []string     // stores the 8 rows for that character
		unescape := func(raw string) string {
			// Convert literal escape sequences to actual characters
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

		// Iterate over each line
		for _, line := range lines {
			if strings.HasPrefix(line, "CHAR:") {
				// If we already have a character being parsed, finalize it
				if currentKey != "" {
					// Pad rows to 8 if needed
					for len(rows) < 8 {
						rows = append(rows, "")
					}
					// Copy rows to banner map
					banner[currentKey] = append([]string(nil), rows...)
				}
				// Set new character key after unescaping
				rawKey := strings.TrimPrefix(line, "CHAR:")
				currentKey = unescape(rawKey)
				rows = nil
				continue
			}

			// Skip lines if no currentKey
			if currentKey == "" {
				continue
			}
			// Append line to current character's rows
			rows = append(rows, strings.TrimRight(line, " "))
		}

		// After finishing all lines, store the last character
		if currentKey != "" {
			for len(rows) < 8 {
				rows = append(rows, "")
			}
			banner[currentKey] = append([]string(nil), rows...)
		}

		// Return the parsed banner map
		return banner, nil
	}

	// If no CHAR: headers, parse blank-line-separated blocks format
	var block []string // temporary storage for current glyph
	char := rune(31)   // start before space (ASCII 32)
	emitBlock := func() {
		if len(block) == 0 {
			return // skip empty blocks
		}
		char++ // increment to next ASCII character
		// Ensure 8 rows for compatibility
		for len(block) < 8 {
			block = append(block, "")
		}
		// Assign block to the banner map using rune as string key
		banner[string(char)] = append([]string(nil), block...)
		block = nil
	}

	// Iterate over all lines
	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			// Blank line = end of current glyph block
			emitBlock()
			continue
		}
		// Add line to current block
		block = append(block, line)
	}
	// Emit any trailing block at the end
	emitBlock()

	// Return the completed banner map
	return banner, nil
}
