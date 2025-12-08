package tests

import (
	"testing"

	"ascii-art/pipeline"
)

func TestLoadBanner_ValidBanner(t *testing.T) {
	banner, err := pipeline.LoadBanner("standard")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	// Έλεγχος για μερικούς βασικούς χαρακτήρες
	requiredChars := []string{"A", "B", " ", "1", "!"}
	for _, c := range requiredChars {
		if _, ok := banner[c]; !ok {
			t.Errorf("expected character %q to be present in banner", c)
		}
	}
}

func TestLoadBanner_InvalidBanner(t *testing.T) {
	_, err := pipeline.LoadBanner("nonexistent")
	if err == nil {
		t.Fatal("expected error for nonexistent banner, got nil")
	}
}

func TestLoadBanner_NewlineCharacter(t *testing.T) {
	banner, err := pipeline.LoadBanner("standard")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if _, ok := banner["\n"]; !ok {
		t.Error("expected newline character '\\n' to be present in banner")
	}
}
func TestLoadBanner_SpecialCharacters(t *testing.T) {
	banner, err := pipeline.LoadBanner("standard")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	specialChars := []string{"@", "#", "$", "%", "^", "&", "*", "(", ")", "-", "_", "+", "=", "{", "}", "[", "]", "|", "\\", ":", ";", "\"", "'", "<", ">", ",", ".", "?", "/"}
	for _, c := range specialChars {
		if _, ok := banner[c]; !ok {
			t.Errorf("expected special character %q to be present in banner", c)
		}
	}
}
func TestLoadBanner_WhitespaceCharacters(t *testing.T) {
	banner, err := pipeline.LoadBanner("standard")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	whitespaceChars := []string{" ", "\t", "\r"}
	for _, c := range whitespaceChars {
		if _, ok := banner[c]; !ok {
			t.Errorf("expected whitespace character %q to be present in banner", c)
		}
	}
}
