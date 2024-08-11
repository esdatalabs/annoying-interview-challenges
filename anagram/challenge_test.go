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
	wrd1, wrd2 string
	valid      bool
}

var suite2 = []allowed{
	{"thisismixedspace", "thisismixedspace", false},
	//Words must be of the same length
	{"givenLength", "short", false},
	//Words must be alphabetical
	{"&es_", "word", false},
	//Valid word pairs
	{"fired", "fried", true},
}

func Test_eligibleWordPairs(t *testing.T) {

	for _, entry := range suite2 {
		if response := validate(entry.wrd1, entry.wrd2); response != entry.valid {
			t.Errorf("Expected %v for validating %s and %s. Got %v", entry.valid, entry.wrd1, entry.wrd2, response)
		}
	}

}
