package tests

// Import the testing package for writing unit tests
import "testing"

// TestReadBanner is a test function that verifies the ReadBanner function works correctly
func TestReadBanner(t *testing.T) {
	// Call ReadBanner with "standard.txt" and capture the returned banner and error
	banner, err := ReadBanner("standard.txt")

	// Check if an error occurred; if so, fail the test with the error message
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Verify that the banner has exactly 95 characters; if not, fail the test with actual length
	if len(banner) != 95 {
		t.Fatalf("expected banner length 95, got %d", len(banner))
	}
}
