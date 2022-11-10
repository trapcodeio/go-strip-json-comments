package stripjsoncomments

import (
	"encoding/json"
	"os"
	"testing"
)

// Package is an exact port of `Sindresorhus` strip-json-comments nodejs package.
// Ported by `trapcodeio`
// Additional test to read from file and compare with expected output

func TestFileSamples(t *testing.T) {
	t.Run("sample.json", func(t *testing.T) {
		testFile("sample.json", t)
	})

	t.Run("sample-big.json", func(t *testing.T) {
		testFile("sample-big.json", t)
	})
}

// ============================================================================
// ============================================================================
// ===== Test Helper Functions ================================================
// ============================================================================
// ============================================================================

// testFile is a shorthand function for testing files
func testFile(filePath string, t *testing.T) {
	// get current directory
	dir, err := os.Getwd()
	if err != nil {
		t.Errorf("Error getting current directory: %v", err)
		return
	}

	// file path
	filePath = dir + "/" + filePath

	// read file
	jsonBytes, err := os.ReadFile(filePath)
	if err != nil {
		t.Errorf("Error opening file: %v", err)
		return
	}

	// we expect an error here since the file contains comments
	if json.Valid(jsonBytes) {
		t.Errorf("Expected json to be invalid")
		return
	}

	// strip comments
	striped := Strip(string(jsonBytes))

	// we expect no error here so if there is an error, we fail the test
	if !json.Valid([]byte(striped)) {
		t.Errorf("Expected no error parsing json without comment, got: %v", err)
		return
	}
}
