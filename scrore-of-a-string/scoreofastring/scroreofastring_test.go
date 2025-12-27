package scoreofastring

import (
	"testing"
)

func TestScoreOfAString(t *testing.T) {
	tests := []struct {
		testName   string
		testString string
		expected   int
		wantErr    bool
	}{
		{
			"Code example",
			"code",
			24,
			false,
		},
		{
			"Kittens rule example",
			"kittens rule",
			226,
			false,
		},
		{
			"Neetcode example",
			"neetcode",
			65,
			false,
		},
		{
			"Empty string",
			"",
			0,
			true,
		},
		{
			"Single character",
			"s",
			0,
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			result, err := ScoreOfAString(tt.testString)

			if tt.wantErr {
				if err == nil {
					t.Errorf("expected error, got nil")
					return
				}

				if err.Error() != "string must contain at least two characters" {
					t.Errorf("unexpected error message: %v", err)
				}
				return
			}

			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if result != tt.expected {
				t.Errorf("want: %v, got: %v", tt.expected, result)
			}
		})
	}
}
