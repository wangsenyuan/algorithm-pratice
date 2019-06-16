package manachers

import "testing"

func runSample(t *testing.T, s string, expect string) {
	res := longestPalindrome(s)
	if res != expect {
		t.Errorf("Sample %s expect %s, but got %s", s, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "bb", "bb")
}
