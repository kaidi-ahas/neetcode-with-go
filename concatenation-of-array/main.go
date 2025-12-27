package main

import (
	"concatenation-of-array/concatenationofarray"
	"fmt"
)

func main() {

	myTests := [][]int{
		{1, 2, 3, 4},
		{},
		{42},
		{-1, -2},
	}

	for _, test := range myTests {
		fmt.Println(concatenationofarray.GetConcatenation(test))
	}
}
