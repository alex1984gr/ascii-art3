package tests

// Import the testing package for writing unit tests and the pipeline package under test
import (
	"testing"

	"ascii-art/pipeline"
)

// TestLoadBanner_ValidBanner checks that loading the "standard" banner succeeds and
// that a few representative characters exist in the returned banner mapping.
func TestLoadBanner_ValidBanner(t *testing.T) {
	// Call LoadBanner with the name "standard" and capture the returned banner map and error
	banner, err := pipeline.LoadBanner("standard")
	// If an error occurred, fail the test immediately with the error message
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	// Define a small list of characters we expect to be present in a typical banner
	requiredChars := []string{"A", "B", " ", "1", "!"}
	// Iterate over each expected character and verify it appears as a key in the banner map
	for _, c := range requiredChars {
		if _, ok := banner[c]; !ok {
			// If a character is missing, report an error but continue checking others
			t.Errorf("expected character %q to be present in banner", c)
		}
	}
}

// TestLoadBanner_InvalidBanner verifies that requesting a nonexistent banner returns an error
func TestLoadBanner_InvalidBanner(t *testing.T) {
	// Call LoadBanner with a name that should not exist and ignore the returned banner
	_, err := pipeline.LoadBanner("nonexistent")
	// Expect an error; if nil, fail the test
	if err == nil {
		t.Fatal("expected error for nonexistent banner, got nil")
	}
}

// TestLoadBanner_NewlineCharacter ensures the newline token is present in the banner map
func TestLoadBanner_NewlineCharacter(t *testing.T) {
	// Load the standard banner again
	banner, err := pipeline.LoadBanner("standard")
	// Fail immediately if loading failed
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	// Check that the newline string "\n" exists as a key in the banner map
	if _, ok := banner["\n"]; !ok {
		// Report an error if newline is missing
		t.Error("expected newline character '\n' to be present in banner")
	}
}

// TestLoadBanner_SpecialCharacters verifies that common special characters are present
func TestLoadBanner_SpecialCharacters(t *testing.T) {
	// Load the standard banner and check for errors
	banner, err := pipeline.LoadBanner("standard")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	// List of special characters that should be represented in the banner mapping
	specialChars := []string{"@", "#", "$", "%", "^", "&", "*", "(", ")", "-", "_", "+", "=", "{", "}", "[", "]", "|", "\\", ":", ";", "\"", "'", "<", ">", ",", ".", "?", "/"}
	// Verify each special character appears as a key in the banner map
	for _, c := range specialChars {
		if _, ok := banner[c]; !ok {
			t.Errorf("expected special character %q to be present in banner", c)
		}
	}
}

// TestLoadBanner_WhitespaceCharacters ensures common whitespace tokens are present
func TestLoadBanner_WhitespaceCharacters(t *testing.T) {
	// Load banner and fail the test if an error occurs
	banner, err := pipeline.LoadBanner("standard")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	// Check for space, tab and carriage return tokens
	whitespaceChars := []string{" ", "\t", "\r"}
	for _, c := range whitespaceChars {
		if _, ok := banner[c]; !ok {
			t.Errorf("expected whitespace character %q to be present in banner", c)
		}
	}
}
