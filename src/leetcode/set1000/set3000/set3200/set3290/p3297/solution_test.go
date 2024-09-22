package p3297

import "testing"

func runSample(t *testing.T, w1 string, w2 string, expect int) {
	res := validSubstringCount(w1, w2)

	if res != int64(expect) {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	word1 := "bcca"
	word2 := "abc"
	expect := 1
	runSample(t, word1, word2, expect)
}

func TestSample2(t *testing.T) {
	word1 := "abcabc"
	word2 := "abc"
	expect := 10
	runSample(t, word1, word2, expect)
}

func TestSample3(t *testing.T) {
	word1 := "abcabc"
	word2 := "aaabc"
	expect := 0
	runSample(t, word1, word2, expect)
}
