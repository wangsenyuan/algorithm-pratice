package p1163

import "testing"

func runSample(t *testing.T, s string, expect string) {
	res := lastSubstring(s)
	if res != expect {
		t.Errorf("Sample %s, expect %s, but got %s", s, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "abab", "bab")
}

func TestSample2(t *testing.T) {
	runSample(t, "leetcode", "tcode")
}

func TestSample3(t *testing.T) {
	runSample(t, "aaaaaaaa", "aaaaaaaa")
}
