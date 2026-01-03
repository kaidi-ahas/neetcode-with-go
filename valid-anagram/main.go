package main

import (
	"fmt"
	"log"
	"valid-anagram/isanagram"
)

func main() {
	ok, err := isanagram.IsAnagram("silent", "listen")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(ok)
}
