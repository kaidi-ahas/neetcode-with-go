package hasduplicate

import (
	"errors"
	"testing"
)

// What behavior do I want to guarantee never breaks in HasDuplicate?
// // The function must reliably detect duplicates
// What does HasDuplicate promise to do?
// // It promises to identify whether a slice has a duplicate.
// What does HasDuplicate return on success?
// // Return true if the slice contains at least one duplicate, false otherwise and nil indicating no error
// What does HasDuplicate return on failure?
// // false boolean and ErrNilSlice or ErrEmptySlice
// Contract:
// // Input: slice of ints
// // Output: boolean, error
// // Failure: empty or nil slice

func TestHasDuplicate(t *testing.T) {

	type Test struct {
		name         string
		testSlice    []int
		expectedBool bool
		wantErr      bool
		expectedErr  error
	}

	tests := []Test{
		{
			"nil slice",
			nil,
			false,
			true,
			ErrNilSlice,
		},
		{
			"empty slice",
			[]int{},
			false,
			true,
			ErrEmptySlice,
		},
		{
			"has duplicate",
			[]int{1, 2, 3, 2},
			true,
			false,
			nil,
		},
		{
			"no duplicate",
			[]int{1, 2, 3},
			false,
			false,
			nil,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			resultBool, err := HasDuplicate(tt.testSlice)

			if tt.wantErr {
				if err == nil {
					t.Fatalf("expected error %v, got nil", tt.expectedErr)
					return
				}
				if !errors.Is(err, tt.expectedErr) {
					t.Fatalf("unexpected error: %v, got: %v", tt.expectedErr, err)
				}
				return
			}

			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if tt.expectedBool != resultBool {
				t.Errorf("Expected: %v, got: %v", tt.expectedBool, resultBool)
			}
		})
	}
}
