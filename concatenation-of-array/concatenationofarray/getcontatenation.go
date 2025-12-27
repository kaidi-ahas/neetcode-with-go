package concatenationofarray

import "errors"

var ErrEmptySlice = errors.New("empty slice")

func GetConcatenation(nums []int) ([]int, error) {

	if len(nums) == 0 {
		return []int{}, ErrEmptySlice
	}

	ans := append([]int{}, nums...) // ans = [nums]

	ans = append(ans, nums...) // ans = [nums*2]

	return ans, nil
}
