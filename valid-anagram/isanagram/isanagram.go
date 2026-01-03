package isanagram

import (
	"errors"
	"unicode"
)

var ErrorEmptyString = "string cannot be empty"
var ErrorUnequalLengths = "string lengths should be equal"
var ErrorUpperCaseDetected = "string cannot have any uppercase letters"
var ErrorWrongCharacter = "only lowercase English letters are allowed"

// Space O(1), Time O(n)
func IsAnagram(s, t string) (error, bool) {

	if hasUpperCaseLetter(s) || hasUpperCaseLetter(t) {
		return errors.New(ErrorUpperCaseDetected), false
	}

	if len(s) == 0 || len(t) == 0 {
		return errors.New(ErrorEmptyString), false

	}

	if len(s) != len(t) {
		return errors.New(ErrorUnequalLengths), false
	}

	if s == t {
		return nil, true
	}

	var freq [26]int // for storing the English alphabet 26 letters

	for i := 0; i < len(s); i++ {
		// only English lowercase letters
		if s[i] < 'a' || s[i] > 'z' || t[i] < 'a' || t[i] > 'z' {
			return errors.New(ErrorWrongCharacter), false
		}

		// increment freq value by 1 at the index where a letter from s is
		freq[s[i]-'a']++

		// decrement freq value by 1 at the index where a letter from t is
		freq[t[i]-'a']--
	}
	return nil, true
}

func hasUpperCaseLetter(str string) bool {
	for _, s := range str {
		if unicode.IsUpper(s) {
			return true
		}
	}
	return false
}
