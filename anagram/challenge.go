package main

import (
	"fmt"
	"strings"
)

func solultion1(w1 string, w2 string) bool {

	howLong := len(w1)
	thisLong := len(w2)

	if howLong != thisLong {
		return false
	}

	return false
}

func validate(w1 string, w2 string) {

}

func normalize(word string) string {

	normalized := strings.ReplaceAll(word, " ", "")

	return normalized
}

func main() {

	w1 := "fireds"
	w2 := "fried"

	if solultion1(w1, w2) {
		fmt.Printf("%s and %s are anagrams of each other", w1, w2)
	} else {
		fmt.Printf("%s and %s are not anagrams of each other", w1, w2)
	}

	println("end")

}
