package main

import (
	"contains-duplicate/hasduplicate"
	"fmt"
)

func main() {
	myTests := [][]int{
		(nil),
		{},
		{1, 2, 3, 2},
		{1, 2, 3},
	}

	for _, myTest := range myTests {
		fmt.Println(hasduplicate.HasDuplicate(myTest))
	}
}
