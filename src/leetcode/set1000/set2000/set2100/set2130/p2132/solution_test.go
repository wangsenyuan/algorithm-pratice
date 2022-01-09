package p2132

import "testing"

func runSample(t *testing.T, words []string, expect int) {
	res := longestPalindrome(words)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	words := []string{"dd", "aa", "bb", "dd", "aa", "dd", "bb", "dd", "aa", "cc", "bb", "cc", "dd", "cc"}
	expect := 22
	runSample(t, words, expect)
}
