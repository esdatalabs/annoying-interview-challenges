package main

import (
	"fmt"
	"regexp"
	"slices"
	"strings"
)

func anagrams(w1 string, w2 string) bool {

	//Convert both string to arrays of 32bit integers (specifically unicode points)
	//Because having a character type like every other grown ass language would go against the grain
	//of golang core competencies
	w1r := []rune(w1)
	w2r := []rune(w2)

	for _, entry := range w1r {
		if !slices.Contains(w2r, entry) {
			return false
		}
	}

	return true
}

func validate(w1 string, w2 string) bool {

	regex := regexp.MustCompile(`^[a-zA-Z]+$`)
	var lettersOnly = regex.MatchString

	//If the words are the same, they're not anagrams
	if w1 == w2 {
		return false
	}

	//words must be of same length
	if len(w1) != len(w2) {
		return false
	}

	//words must be composed of letters only
	if !lettersOnly(w1) || !lettersOnly(w2) {
		return false
	}

	return true
}

// Ensure we're comparing apples to apples
func normalize(word string) string {

	//Purse all spaces
	normalized := strings.ReplaceAll(word, " ", "")

	//make everything lowercase
	normalized = strings.ToLower(normalized)

	return normalized
}

func main() {

	w1 := "fired"
	w2 := "fried"

	w1n := normalize(w1)
	w2n := normalize(w2)

	eligible := validate(w1n, w2n)

	if eligible && anagrams(w1, w2) {
		fmt.Printf("%s and %s are anagrams of each other", w1, w2)
	} else {
		fmt.Printf("%s and %s are not anagrams of each other", w1, w2)
	}
	println("")
	println("end")

}
