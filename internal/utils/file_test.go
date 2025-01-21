package utils

import (
	"os"
	"path/filepath"
	"testing"
)

// TestLoadNumbersFile validates the proper reading and parsing of a file containing integers.
func TestLoadNumbersFile(t *testing.T) {
	createTempFile := func(content string) (string, func(), error) {
		tmpFile, err := os.CreateTemp("", "test-numbers-*.txt")
		if err != nil {
			return "", nil, err
		}
		_, err = tmpFile.WriteString(content)
		if err != nil {
			return "", nil, err
		}
		cleanup := func() {
			os.Remove(tmpFile.Name())
		}
		return tmpFile.Name(), cleanup, nil
	}

	t.Run("Valid file with numbers", func(t *testing.T) {
		content := "10\n20\n30\n40\n50\n"
		filePath, cleanup, err := createTempFile(content)
		if err != nil {
			t.Fatalf("Failed to create temp file: %v", err)
		}
		defer cleanup()

		numbers, err := LoadNumbersFile(filePath)
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		expected := []int{10, 20, 30, 40, 50}
		for i, num := range numbers {
			if num != expected[i] {
				t.Errorf("Expected %d at index %d, got %d", expected[i], i, num)
			}
		}
	})

	t.Run("File does not exist", func(t *testing.T) {
		invalidFilePath := filepath.Join(os.TempDir(), "nonexistent-file.txt")
		_, err := LoadNumbersFile(invalidFilePath)
		if err == nil {
			t.Fatal("Expected an error for missing file, but got none")
		}
		expectedError := "failed to open file"
		if err != nil && !contains(err.Error(), expectedError) {
			t.Errorf("Expected error containing '%s', got '%v'", expectedError, err)
		}
	})

	t.Run("Malformed content", func(t *testing.T) {
		content := "10\n20\nabc\n40\n"
		filePath, cleanup, err := createTempFile(content)
		if err != nil {
			t.Fatalf("Failed to create temp file: %v", err)
		}
		defer cleanup()

		numbers, err := LoadNumbersFile(filePath)
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		expected := []int{10, 20}
		if len(numbers) != len(expected) {
			t.Fatalf("Expected %d numbers, got %d", len(expected), len(numbers))
		}
		for i, num := range numbers {
			if num != expected[i] {
				t.Errorf("Expected %d at index %d, got %d", expected[i], i, num)
			}
		}
	})
}

// contains checks if a substring exists in a string
func contains(s, substr string) bool {
	return len(substr) == 0 || len(s) >= len(substr) && s[:len(substr)] == substr || len(s) > len(substr) && contains(s[1:], substr)
}
