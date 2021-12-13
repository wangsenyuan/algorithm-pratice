package p2014

import "testing"

func runSample(t *testing.T, s string, k int, expect string) {
	res := longestSubsequenceRepeatedK(s, k)

	if res != expect {
		t.Errorf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "letsleetcode"
	k := 2
	expect := "let"
	runSample(t, s, k, expect)
}

func TestSample2(t *testing.T) {
	s := "bbabbabbbbabaababab"
	k := 3
	expect := "bbbb"
	runSample(t, s, k, expect)
}
