package p730

import "testing"

func TestSample1(t *testing.T) {
	S := "bccb"
	res := countPalindromicSubsequences(S)
	if res != 6 {
		t.Errorf("%s should give 6, but get %d", S, res)
	}
}

func TestSample2(t *testing.T) {
	S := "abcdabcdabcdabcdabcdabcdabcdabcddcbadcbadcbadcbadcbadcbadcbadcba"
	res := countPalindromicSubsequences(S)
	if res != 104860361 {
		t.Errorf("%s should give 104860361, but get %d", S, res)
	}
}
