package stripjsoncomments

import (
	"os"
	"testing"
)

// Package is an exact port of `Sindresorhus` strip-json-comments nodejs package.
// Ported by `trapcodeio`
// File benchmark.js in repository: https://github.com/sindresorhus/strip-json-comments

func BenchmarkWithFiles(b *testing.B) {
	// Get the current directory
	var dir, err = os.Getwd()
	if err != nil {
		b.Errorf("Error getting current directory: %v", err)
		return
	}
	// Get json files in current directory
	json, err := os.ReadFile(dir + "/sample.json")
	bigJson, err := os.ReadFile(dir + "/sample-big.json")
	if err != nil {
		b.Errorf("Error opening required files: %v", err)
		return
	}

	b.Run("strip JSON comments", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Strip(string(json))
		}
	})

	b.Run("strip JSON comments without whitespace", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			StripWithOptions(string(json), &Options{Whitespace: false})
		}
	})

	b.Run("strip Big JSON comments", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Strip(string(bigJson))
		}
	})

}
