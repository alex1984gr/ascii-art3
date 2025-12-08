package pipeline

import (
	"errors"
	"os"
)

// ReadInput reads the contents of the file located at the given path and
// returns the file contents as a string. If an error occurs while reading the
// file, it returns a non-nil error.
func ReadInput(path string) (string, error) {
	// If the caller passed an empty path, return an error immediately.
	if path == "" {
		return "", errors.New("path is empty")
	}

	// Read the entire file into memory. This uses Go 1.16+ os.ReadFile which
	// returns the raw bytes and any error encountered while opening/reading.
	b, err := os.ReadFile(path)
	if err != nil {
		// Propagate the read error to the caller for handling.
		return "", err
	}

	// Convert the bytes to a string and return with a nil error.
	return string(b), nil
}
