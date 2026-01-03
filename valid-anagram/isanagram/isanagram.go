package isanagram

import (
	"errors"
)

var (
	ErrEmptyString      = errors.New("string cannot be empty")
	ErrUnequalLengths   = errors.New("string lengths should be equal")
	ErrInvalidCharacter = errors.New("only lowercase English letters are allowed")
)

// IsAnagram reports whether s and t are anagrams.
// It returns an error if the input is invalid.
func IsAnagram(s, t string) (bool, error) {

	if err := validateInput(s, t); err != nil {
		return false, err
	}

	var freq [26]int

	for i := 0; i < len(s); i++ {
		freq[s[i]-'a']++
		freq[t[i]-'a']--
	}

	for _, v := range freq {
		if v != 0 {
			return false, nil
		}
	}

	return true, nil
}

func validateInput(s, t string) error {
	if len(s) == 0 || len(t) == 0 {
		return ErrEmptyString
	}

	if len(s) != len(t) {
		return ErrUnequalLengths
	}

	for i := 0; i < len(s); i++ {
		if !isLowercaseASCII(s[i]) || !isLowercaseASCII(t[i]) {
			return ErrInvalidCharacter
		}
	}
	return nil
}

func isLowercaseASCII(b byte) bool {
	return b >= 'a' && b <= 'z'
}
