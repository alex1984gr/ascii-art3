package tests

import (
	"os"
	"testing"

	"ascii-art/pipeline"
)

func TestReadInput_FileExists(t *testing.T) {
	// Create temporary file
	tmp := "test_input.txt"
	content := "Hello"
	err := os.WriteFile(tmp, []byte(content), 0644)
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}
	defer os.Remove(tmp)

	// Run function
	result, err := pipeline.ReadInput(tmp)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if result != content {
		t.Fatalf("expected %q, got %q", content, result)
	}
}

func TestReadInput_FileNotFound(t *testing.T) {
	_, err := pipeline.ReadInput("no_such_file.txt")
	if err == nil {
		t.Fatalf("expected error for missing file, got nil")
	}
}

func TestReadInput_EmptyFile(t *testing.T) {
	tmp := "empty.txt"
	os.WriteFile(tmp, []byte(""), 0644)
	defer os.Remove(tmp)

	result, err := pipeline.ReadInput(tmp)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	if result != "" {
		t.Fatalf("expected empty string, got: %q", result)
	}
}
