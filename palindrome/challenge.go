package main

import (
	"fmt"
	"regexp"
	"slices"
	"strings"
)

func palindrome(phrase string) bool {

	size := len(phrase)

	//Convert to character array
	rphrase := make([]rune, size)
	for index, ent := range phrase {
		rphrase[index] = ent
	}

	//Copy & Reverse
	rphraserev := make([]rune, size)
	copy(rphraserev, rphrase)
	slices.Reverse(rphraserev)

	//Get halfway mark of slice
	ceiling := len(rphrase) / 2

	//If the the slide is oddly sized, add an offset
	if len(rphrase)%2 != 0 {
		ceiling += 1
	}

	//Iterate through "half" the slices. If the contents dont match, its not a palindrome
	for i := 0; i < ceiling; i++ {
		if rphrase[i] != rphraserev[i] {
			return false
		}
	}

	return true
}

func validate(phrase string) bool {

	var lettersOnly = regexp.MustCompile(`^[a-zA-Z]+$`).MatchString

	//Words cannot be empty
	if len(phrase) == 0 {
		return false
	}

	//words must be composed of letters only
	return lettersOnly(phrase)
}

// Ensure we're comparing apples to apples
func normalize(word string) string {

	//replace non letter
	reg, _ := regexp.Compile(`[^a-zA-Z0-9 ]+`)
	normalized := reg.ReplaceAllString(word, "")

	//Purse all spaces
	normalized = strings.ReplaceAll(normalized, " ", "")

	//make everything lowercase
	return strings.ToLower(normalized)
}

func main() {

	phrase := "racecar"

	phraseNormalized := normalize(phrase)

	eligible := validate(phraseNormalized)

	if eligible && palindrome(phraseNormalized) {
		fmt.Printf("%s is a palindrome", phrase)
	} else {
		fmt.Printf("%s is not a palindrome", phrase)
	}
	println("")
	println("end")

}
