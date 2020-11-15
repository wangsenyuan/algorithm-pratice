package p5603

import "testing"

func runSample(t *testing.T, x, y string, expect bool) {
	res := closeStrings(x, y)

	if res != expect {
		t.Errorf("Sample %s %s, expect %t, but got %t", x, y, expect, res)
	}
}

func TestSample1(t *testing.T) {
	word1 := "abc"
	word2 := "bca"
	expect := true
	runSample(t, word1, word2, expect)
}

func TestSample2(t *testing.T) {
	word1 := "a"
	word2 := "aa"
	expect := false
	runSample(t, word1, word2, expect)
}

func TestSample3(t *testing.T) {
	word1 := "cabbba"
	word2 := "abbccc"
	expect := true
	runSample(t, word1, word2, expect)
}

func TestSample4(t *testing.T) {
	word1 := "cabbba"
	word2 := "aabbss"
	expect := false
	runSample(t, word1, word2, expect)
}

func TestSample5(t *testing.T) {
	word1 := "cabbba"
	word2 := "aabbbs"
	expect := false
	runSample(t, word1, word2, expect)
}

func TestSample6(t *testing.T) {
	word1 := "cabbba"
	word2 := "ccbbba"
	expect := true
	runSample(t, word1, word2, expect)
}
