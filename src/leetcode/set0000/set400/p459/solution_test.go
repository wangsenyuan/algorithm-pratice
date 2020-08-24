package p459

import "testing"

func runSample(t *testing.T, s string, expect bool) {
	res := repeatedSubstringPattern(s)

	if res != expect {
		t.Errorf("Sample %s, expect %t, but got %t", s, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "abab", true)
}

func TestSample2(t *testing.T) {
	runSample(t, "abb", false)
}

func TestSample3(t *testing.T) {
	runSample(t, "abcabcabcabc", true)
}
