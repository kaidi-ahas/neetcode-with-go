package concatenationofarray

import (
	"errors"
	"reflect"
	"testing"
)

// What behavior do I want to guarantee never breaks in func GetConcatenation?
// // What does this function promise to do?
// // What does it return on success?
// // What does it return on failure?
// // // Contract:
// // // Input: slice of ints
// // // Output: slice of ints that contains a duplication of input slice of ints, error
// // // Failure: empty or nil slice
func TestGetConcatenation(t *testing.T) {
	tests := []struct {
		name      string
		testSlice []int
		expected  []int
		wantErr   bool
	}{
		{
			"Regular slice",
			[]int{1, 2, 3, 4},
			[]int{1, 2, 3, 4, 1, 2, 3, 4},
			false,
		},
		{
			"Empty slice",
			[]int{},
			[]int{},
			true,
		},
		{
			"Nil slice",
			nil,
			[]int{},
			true,
		},
		{
			"Single element",
			[]int{42},
			[]int{42, 42},
			false,
		},
		{
			"Negative numbers",
			[]int{-1, -2},
			[]int{-1, -2, -1, -2},
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt := tt // capture the range variable
			result, err := GetConcatenation(tt.testSlice)

			if tt.wantErr {
				if err == nil {
					t.Errorf("Expected error, got nil")
					return
				}

				if !errors.Is(err, ErrEmptySlice) {
					t.Errorf("unexpected error message: %v", err)
				}
				return
			}

			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Want: %v, got: %v", tt.expected, result)
			}

		})
	}

}
