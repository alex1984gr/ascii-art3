package tests

// Import the os package for file operations, testing for unit tests, and the pipeline module
import (
	"os"
	"testing"

	"ascii-art/pipeline"
)

// TestReadInput_FileExists verifies that ReadInput correctly reads from an existing file
func TestReadInput_FileExists(t *testing.T) {
	// Define the temporary file name to be created
	tmp := "test_input.txt"
	// Define the content that will be written to the file
	content := "Hello"
	// Write the content to the temporary file with read/write permissions (0644)
	err := os.WriteFile(tmp, []byte(content), 0644)
	// Check if file creation failed and terminate the test if it did
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}
	// Schedule the temporary file to be deleted after the test completes
	defer os.Remove(tmp)

	// Call ReadInput with the temporary file and capture the result and any error
	result, err := pipeline.ReadInput(tmp)
	// Check if an error occurred; if so, fail the test
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	// Verify that the returned content matches the original content; if not, fail the test
	if result != content {
		t.Fatalf("expected %q, got %q", content, result)
	}
}

// TestReadInput_FileNotFound verifies that ReadInput returns an error when the file doesn't exist
func TestReadInput_FileNotFound(t *testing.T) {
	// Call ReadInput with a non-existent file name and capture the error
	_, err := pipeline.ReadInput("no_such_file.txt")
	// Verify that an error was returned; if not, fail the test
	if err == nil {
		t.Fatalf("expected error for missing file, got nil")
	}
}

// TestReadInput_EmptyFile verifies that ReadInput correctly handles empty files
func TestReadInput_EmptyFile(t *testing.T) {
	// Define the temporary file name to be created
	tmp := "empty.txt"
	// Write an empty string to the temporary file
	os.WriteFile(tmp, []byte(""), 0644)
	// Schedule the temporary file to be deleted after the test completes
	defer os.Remove(tmp)

	// Call ReadInput with the empty file and capture the result and any error
	result, err := pipeline.ReadInput(tmp)
	// Check if an error occurred; if so, fail the test
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	// Verify that the result is an empty string; if not, fail the test
	if result != "" {
		t.Fatalf("expected empty string, got: %q", result)
	}
}
