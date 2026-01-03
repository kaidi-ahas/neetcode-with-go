package isanagram

import (
	"errors"
	"testing"
)

// What behavior do I want to guarantee never breaks in IsAnagram?
// // The function must reliably detect whether two strings are anagrams.
// What does IsAnagram return on success?
// // Returns true if the two strings are anagrams.
// What does IsAnagram return on failure?
// // error, false
// Contract:
// // Input: two strings
// // Output: boolean, error
// // Failure: error, false
func TestIsAnagram(t *testing.T) {
	type Test struct {
		name    string
		s, t    string
		wantErr error
		want    bool
	}

	tests := []Test{
		{
			"basic anagram",
			"silent",
			"listen",
			nil,
			true,
		},
		{
			"single character",
			"m",
			"m",
			nil,
			true,
		},
		{
			"empty string",
			"",
			"hi",
			ErrEmptyString,
			false,
		},
		{
			"unequal lengths",
			"crazy",
			"craz",
			ErrUnequalLengths,
			false,
		},
		{
			"contains numbers",
			"stuffy123",
			"stuffy123",
			ErrInvalidCharacter,
			false,
		},
		{
			"contains uppercase",
			"Me",
			"Me",
			ErrInvalidCharacter,
			false,
		},
		{
			"all uppercase",
			"YAY",
			"YAY",
			ErrInvalidCharacter,
			false,
		},
		{
			"invalid symbol",
			"hello?",
			"hello?",
			ErrInvalidCharacter,
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := IsAnagram(tt.t, tt.s)

			if tt.wantErr != nil {
				if err == nil {
					t.Fatalf("expected error: %v, got nil", tt.wantErr)
				}
				if !errors.Is(err, tt.wantErr) {
					t.Fatalf("unexpected error: %v message: %v", tt.wantErr, err)
				}
				return
			}

			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if tt.want != got {
				t.Errorf("want: %v, got: %v", tt.want, got)
			}
		})
	}
}
