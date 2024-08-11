package main

func solultion1(w1 string, w2 string) bool {

	howLong := len(w1)
	thisLong := len(w2)

	//The
	if howLong != thisLong {
		return false
	}

	return false
}

func main() {

	w1 := "fireds"
	w2 := "fried"

	if solultion1(w1, w2) {
		println("%s and %s are anagrams of each other", w2, w2)
	} else {
		println("%s and %s are not anagrams of each other", w2, w2)
	}

}
