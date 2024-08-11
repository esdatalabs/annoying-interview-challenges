package main

import (
	"regexp"
	"strings"
)

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

}
