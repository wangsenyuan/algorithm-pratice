package p2531

import "testing"

func runSample(t *testing.T, word1, word2 string, expect bool) {
	res := isItPossible(word1, word2)

	if res != expect {
		t.Errorf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	w1 := "ac"
	w2 := "b"
	expect := false
	runSample(t, w1, w2, expect)
}

func TestSample2(t *testing.T) {
	w1 := "aa"
	w2 := "ab"
	expect := false
	runSample(t, w1, w2, expect)
}
