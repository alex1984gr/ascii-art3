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
	// Validate that the banner name is not empty.
	if name == "" {
		return nil, errors.New("empty banner name")
	}

	// Construct the file path by joining the "banners" directory with the name + ".txt" extension.
	// Start from the current working directory. If tests run from the workspace root, this works directly.
	// If tests run from tests/, adjust the path to use ../banners.
	path := filepath.Join("banners", name+".txt")

	// Attempt to open the banner file at the initially constructed path.
	f, err := os.Open(path)
	if err != nil {
		// If the file was not found at the first path, try with "../banners" prefix (tests running from tests/).
		path = filepath.Join("..", "banners", name+".txt")
		var err2 error
		f, err2 = os.Open(path)
		if err2 != nil {
			// If both paths failed, return the original error.
			return nil, fmt.Errorf("open banner file: %w", err)
		}
	}
	// Ensure the file is closed when the function returns.
	defer f.Close()

	// Parse the opened file using the common reader-based parser.
	return LoadBannerFromReader(f)
}

// LoadBannerFromReader parses a banner font from an io.Reader in compact format.
// The format is:
//
//	CHAR:<character or escape sequence>
//	<row 1>
//	<row 2>
//	...
//	<row 8>
//
// Escape sequences like "\n", "\t", "\r" are unescaped to their actual character values.
// The function returns a map with character keys and 8-row glyph values.
func LoadBannerFromReader(r io.Reader) (map[string][]string, error) {
	// Create a scanner to read the input line by line.
	scanner := bufio.NewScanner(r)

	// Initialize the banner map to store glyph definitions.
	banner := make(map[string][]string)

	// Track the current character key being processed.
	var currentKey string

	// Track the rows (lines) being accumulated for the current character.
	var rows []string

	// Helper function to unescape special character sequences (e.g., "\n" → actual newline).
	unescape := func(raw string) string {
		// Replace the escaped sequence with the actual character value.
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

	// Scan through each line of the input file.
	for scanner.Scan() {
		// Get the current line as a string.
		line := scanner.Text()

		// If this line starts with "CHAR:", it marks the beginning of a new glyph block.
		if strings.HasPrefix(line, "CHAR:") {
			// If we were already processing a previous glyph, save it first.
			if currentKey != "" {
				// Ensure the previous glyph has exactly 8 rows (pad with empty strings if needed).
				for len(rows) < 8 {
					rows = append(rows, "")
				}
				// Create a copy of the rows and store it in the banner map.
				banner[currentKey] = append([]string(nil), rows...)
			}

			// Extract the character sequence after "CHAR:" and unescape it.
			rawKey := strings.TrimPrefix(line, "CHAR:")
			currentKey = unescape(rawKey)

			// Reset the rows slice for the new glyph.
			rows = nil
			continue
		}

		// If we have not yet encountered a "CHAR:" line, skip this line (e.g., comments, blank lines).
		if currentKey == "" {
			continue
		}

		// Add this line as a row to the current glyph being built.
		rows = append(rows, line)
	}

	// After scanning all lines, save the last glyph if one was being processed.
	if currentKey != "" {
		// Ensure the last glyph has exactly 8 rows (pad with empty strings if needed).
		for len(rows) < 8 {
			rows = append(rows, "")
		}
		// Create a copy and store it in the banner map.
		banner[currentKey] = append([]string(nil), rows...)
	}

	// Check if the scanner encountered any errors during reading.
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	// Return the fully populated banner map.
	return banner, nil
}
