package p2063

import "testing"

func runSample(t *testing.T, word string, expect int) {
	res := int(countVowels(word))

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	word := "noosabasboosa"
	expect := 237
	runSample(t, word, expect)
}
