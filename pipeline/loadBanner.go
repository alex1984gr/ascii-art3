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
	// Read all lines first so we can detect the file's format and then parse accordingly.
	scanner := bufio.NewScanner(r)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	// Quick format detection: if any line starts with "CHAR:", treat this as the
	// compact CHAR:<literal> format. Otherwise treat it as the common
	// blank-line-separated glyph blocks format used by some remote repos.
	usesCharHeader := false
	for _, l := range lines {
		if strings.HasPrefix(l, "CHAR:") {
			usesCharHeader = true
			break
		}
	}

	banner := make(map[string][]string)

	if usesCharHeader {
		// Parse the existing CHAR: format (backwards compatible with earlier behavior).
		var currentKey string
		var rows []string
		unescape := func(raw string) string {
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

		for _, line := range lines {
			if strings.HasPrefix(line, "CHAR:") {
				if currentKey != "" {
					for len(rows) < 8 {
						rows = append(rows, "")
					}
					banner[currentKey] = append([]string(nil), rows...)
				}
				rawKey := strings.TrimPrefix(line, "CHAR:")
				currentKey = unescape(rawKey)
				rows = nil
				continue
			}

			if currentKey == "" {
				continue
			}
			rows = append(rows, line)
		}

		if currentKey != "" {
			for len(rows) < 8 {
				rows = append(rows, "")
			}
			banner[currentKey] = append([]string(nil), rows...)
		}

		return banner, nil
	}

	// Parse blank-line-separated blocks format. This format contains glyphs in
	// sequence; we'll map them to ASCII codepoints by incrementing a rune
	// counter. This mirrors how other parts of the remote project map glyphs.
	var block []string
	// Start before space so the first increment maps to 32 (space) if the file
	// is arranged in the standard printable ASCII order.
	char := rune(31)

	emitBlock := func() {
		if len(block) == 0 {
			return
		}
		// Advance to the next character code and assign these rows.
		char++
		// Ensure exactly 8 rows for compatibility with the renderer.
		for len(block) < 8 {
			block = append(block, "")
		}
		// Convert rune to string key (single-character string).
		banner[string(char)] = append([]string(nil), block...)
		block = nil
	}

	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			// blank line => end of current glyph block
			emitBlock()
			continue
		}
		block = append(block, line)
	}
	// emit any trailing block
	emitBlock()

	// Ensure common control tokens exist in the banner map so tests and
	// rendering logic can look them up directly. Use 8 empty rows for each
	// control token if the font file didn't provide them explicitly.
	ensure := func(key string) {
		if _, ok := banner[key]; !ok {
			rows := make([]string, 8)
			for i := range rows {
				rows[i] = ""
			}
			banner[key] = rows
		}
	}

	ensure("\n")
	ensure("\t")
	ensure("\r")

	return banner, nil
}
