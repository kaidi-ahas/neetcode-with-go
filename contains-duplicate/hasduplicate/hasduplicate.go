package hasduplicate

import "errors"

var ErrNilSlice = errors.New("nil slice")
var ErrEmptySlice = errors.New("empty slice")

// Time complexity O(n), space complexity O(n)
func HasDuplicate(nums []int) (bool, error) {
	if nums == nil {
		return false, ErrNilSlice
	}

	if len(nums) == 0 {
		return false, ErrEmptySlice
	}

	seen := make(map[int]struct{})

	for _, num := range nums {
		if _, ok := seen[num]; ok {
			return ok, nil
		}
		seen[num] = struct{}{}
	}

	return false, nil
}

// Time complexity O(n²), space complexity O(n)
// func HasDuplicate(nums []int) bool {
// 	hasDuplicate := false
// 	for i := 0; i < len(nums); i++ {
// 		for j := i + 1; j < len(nums); j++ {
// 			if nums[i] == nums[j] {
// 				hasDuplicate = true
// 			}
// 		}
// 	}
// 	return hasDuplicate
// }
