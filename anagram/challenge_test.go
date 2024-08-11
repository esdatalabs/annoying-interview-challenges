package main

import "testing"

type data struct {
	wrd, normalized string
}

var suite = []data{
	data{"this has spaces", "thishasspaces"},
}

func Test_normalize(t *testing.T) {

	for _, entry := range suite {

		if response := normalize(entry.wrd); response != entry.normalized {
			t.Errorf("Normalized version of %s is %s, got %s instead", entry.wrd, entry.normalized, response)
		}

	}

}
