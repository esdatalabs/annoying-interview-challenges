package main

import (
	"regexp"
	"strings"
)

func validate(phrase string) bool {

	regex := regexp.MustCompile(`^[a-zA-Z]+$`)
	var lettersOnly = regex.MatchString

	//Words cannot be empty
	if len(phrase) == 0 {
		return false
	}

	//words must be composed of letters only
	if !lettersOnly(phrase) {
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
