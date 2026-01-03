package main

import (
	"fmt"
	"valid-anagram/isanagram"
)

func main() {
	myTests := [][]string{
		{"", "hello"},
		{"hi", "Hi"},
		{"t", "T"},
		{"ih", "hi"},
		{"thiSisSoDif", "thiSisSoDif"},
		{"", ""},
		{"cats", "hamsters"},
		{"bye", "byee"},
	}

	for _, v := range myTests {
		fmt.Println(isanagram.IsAnagram(v[0], v[1]))
		fmt.Println()
	}
}
