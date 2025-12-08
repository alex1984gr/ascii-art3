package tests

// Import the testing package for unit tests and the pipeline package under test
import (
	"testing"

	"ascii-art/pipeline"
)

// TestAssembleArt_SingleLine verifies that AssembleArt correctly joins a single character's lines with newlines
func TestAssembleArt_SingleLine(t *testing.T) {
	// Define a slice containing the 8 ASCII art lines representing one character (e.g., 'A')
	lines := []string{
		" _ ",  // Line 1 of the character
		"/ \\", // Line 2 of the character (backslash escaped)
		"|_|",  // Line 3 of the character
		"   ",  // Line 4 (empty row)
		"   ",  // Line 5 (empty row)
		"   ",  // Line 6 (empty row)
		"   ",  // Line 7 (empty row)
		"   ",  // Line 8 (empty row)
	}

	// Call AssembleArt with the lines and capture the joined result
	assembled := pipeline.AssembleArt(lines)
	// [DEBUG] Line assembly execution point — check if AssembleArt correctly joins lines with newlines

	// Define the expected output: all lines joined together with newline characters between them
	expected := " _ \n/ \\\n|_|\n   \n   \n   \n   \n   "
	// Verify that the assembled output matches the expected string
	if assembled != expected {
		// [DEBUG] Assembly verification point — check if output matches expected newline-separated format
		// Report an error with both the expected and actual output in a readable format
		t.Errorf("Expected:\n%q\nGot:\n%q", expected, assembled)
	}
}

// TestAssembleArt_MultipleLines verifies that AssembleArt correctly joins multiple characters' lines with newlines
func TestAssembleArt_MultipleLines(t *testing.T) {
	// Define a slice containing 8 lines of ASCII art for two concatenated characters (e.g., 'AB')
	lines := []string{
		" _ __ ",  // Line 1 for both characters
		"/ \\|_)", // Line 2 for both characters
		"|_||_)",  // Line 3 for both characters
		"      ",  // Line 4 (all spaces)
		"      ",  // Line 5 (all spaces)
		"      ",  // Line 6 (all spaces)
		"      ",  // Line 7 (all spaces)
		"      ",  // Line 8 (all spaces)
	}

	// Call AssembleArt to join the multi-character lines
	assembled := pipeline.AssembleArt(lines)
	// [DEBUG] Multi-line assembly point — check if AssembleArt preserves line content integrity

	// Define the expected output: lines joined with newline separators
	expected := " _ __ \n/ \\|_)\n|_||_)\n      \n      \n      \n      \n      "
	// Verify the result matches the expected joined string
	if assembled != expected {
		// Report error showing both expected and actual outputs
		t.Errorf("Expected:\n%q\nGot:\n%q", expected, assembled)
	}
}

// TestAssembleArt_EmptyLines verifies that AssembleArt correctly handles an empty input slice
func TestAssembleArt_EmptyLines(t *testing.T) {
	// Define an empty slice with no lines to assemble
	lines := []string{}

	// Call AssembleArt with the empty slice
	assembled := pipeline.AssembleArt(lines)

	// Verify that the result is an empty string (no content to join)
	if assembled != "" {
		// Report an error if the output is not empty
		t.Errorf("Expected empty string, got %q", assembled)
	}
}

// TestAssembleArt_WithTrailingNewline verifies that AssembleArt correctly joins lines without adding a trailing newline
func TestAssembleArt_WithTrailingNewline(t *testing.T) {
	// Define a slice with only 3 lines (fewer than the typical 8-row banner format)
	lines := []string{
		" _ ",  // First line
		"/ \\", // Second line
		"|_|",  // Third line
	}

	// Call AssembleArt to join the lines
	assembled := pipeline.AssembleArt(lines)

	// Define the expected output: lines joined with newlines, NO trailing newline at the end
	expected := " _ \n/ \\\n|_|"
	// Verify that the output matches and contains no unwanted trailing newline
	if assembled != expected {
		// Report error with formatted output for debugging
		t.Errorf("Expected:\n%q\nGot:\n%q", expected, assembled)
	}
}
