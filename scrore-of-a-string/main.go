package main

import (
	"fmt"
	"score-of-a-string/scoreofastring"
)

func main() {

	words := []string{"code", "kittens rule", "neetcode", "i", ""}

	for _, word := range words {
		fmt.Println(scoreofastring.ScoreOfAString(word))
	}
}
