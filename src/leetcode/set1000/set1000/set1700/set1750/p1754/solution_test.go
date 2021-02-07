package p1754

import "testing"

func runSample(t *testing.T, word1, word2 string, expect string) {
	res := largestMerge(word1, word2)

	if res != expect {
		t.Errorf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	word1 := "cabaa"
	word2 := "bcaaa"
	expect := "cbcabaaaaa"
	runSample(t, word1, word2, expect)
}

func TestSample2(t *testing.T) {
	word1 := "abcabc"
	word2 := "abdcaba"
	expect := "abdcabcabcaba"
	runSample(t, word1, word2, expect)
}
