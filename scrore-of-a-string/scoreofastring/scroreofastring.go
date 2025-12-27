package scoreofastring

import (
	"errors"
	"math"
)

// Space complexity O(1)
// Time complexity O(n)
func ScoreOfAString(s string) (int, error) {
	if len(s) < 2 {
		return 0, errors.New("string must contain at least two characters")
	}

	var sum int

	for i := 0; i < len(s)-1; i++ {
		sum += int(math.Abs(float64(s[i+1]) - float64(s[i])))
	}
	return sum, nil
}
