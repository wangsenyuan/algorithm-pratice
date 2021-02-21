package p1770

import "testing"

func runSample(t *testing.T, word1, word2 string, expect int) {
	res := longestPalindrome(word1, word2)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	word1 := "cacb"
	word2 := "cbba"
	expect := 5
	runSample(t, word1, word2, expect)
}

func TestSample2(t *testing.T) {
	word1 := "ab"
	word2 := "ab"
	expect := 3
	runSample(t, word1, word2, expect)
}

func TestSample3(t *testing.T) {
	word1 := "aa"
	word2 := "bb"
	expect := 0
	runSample(t, word1, word2, expect)
}

func TestSample4(t *testing.T) {
	word1 := "ceebeddc"
	word2 := "d"
	expect := 3
	runSample(t, word1, word2, expect)
}
