package main

import "testing"

type data struct {
	wrd, normalized string
}

var suite = []data{
	//Normalizing should prune spaces
	{"this has spaces", "thishasspaces"},
	//Normalzing should convert word to lower case
	{"ThisISMIXEDSPACE", "thisismixedspace"},
}

func Test_normalize(t *testing.T) {

	for _, entry := range suite {

		if response := normalize(entry.wrd); response != entry.normalized {
			t.Errorf("Normalized version of %s is %s, got %s instead", entry.wrd, entry.normalized, response)
		}

	}

}

type allowed struct {
	wrd   string
	valid bool
}

var suite2 = []allowed{
	{"thisismixedspace", true},
	//Words cannot be empty
	{"", false},
	//Words must be alphabetical
	{"&es_", false},
	//Valid palindrome
	{"Mr. Owl ate my metal worm", true},
}

func Test_eligibleWordPairs(t *testing.T) {

	for _, entry := range suite2 {
		if response := validate(entry.wrd); response != entry.valid {
			t.Errorf("Expected %v for validating %s. Got %v", entry.valid, entry.wrd, response)
		}
	}

}
