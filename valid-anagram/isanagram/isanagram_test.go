package isanagram

import (
	"testing"
)

// What behavior do I want to guarantee never breaks in hasUpperCaseLetter?
// // The function must reliably detect an uppercase letter
// What does hasUpperCaseLetter return on success?
// // Returns true if the string contains at least one uppercase letter, false otherwise
// What does hasUpperCaseLetter return on failure?
// // haven't implemented this
// Contract:
// // Input: a string
// // Output: boolean
// // Failure: not implemented
func TestHasUpperCaseLetter(t *testing.T) {
	type Test struct {
		name  string
		input string
		want  bool
	}

	tests := []Test{
		{
			"has uppercase letter",
			"Dracula",
			true,
		},
		{
			"no uppercase letter",
			"morning",
			false,
		},
		{
			"all uppercase letters",
			"RUN",
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := hasUpperCaseLetter(tt.input)

			if tt.want != result {
				t.Errorf("want %v, got %v", tt.want, result)
				return
			}
		})
	}
}

// What behavior do I want to guarantee never breaks in IsAnagram?
// // The function must reliably detect whether two strings are anagrams.
// What does IsAnagram return on success?
// // Returns true if the two strings are anagrams.
// What does IsAnagram return on failure?
// // false boolean and ErrNilSlice or ErrEmptySlice
// Contract:
// // Input: slice of ints
// // Output: boolean, error
// // Failure: empty or nil slice
func TestIsAnagram(t *testing.T) {
	type Test struct {
		name    string
		input   string
		wantErr bool
	}
}
