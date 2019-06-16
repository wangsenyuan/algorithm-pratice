package p1092

import "testing"

func runSample(t *testing.T, s1, s2 string, expect string) {
	res := shortestCommonSupersequence(s1, s2)
	if len(res) != len(expect) {
		t.Fatalf("%s %s expect %s, but got %s, different len", s1, s2, expect, res)
	}

	if !isSubsequence(res, s1) {
		t.Fatalf("%s %s expect %s, but got %s, is not a super seq of s1", s1, s2, expect, res)
	}

	if !isSubsequence(res, s2) {
		t.Fatalf("%s %s expect %s, but got %s, is not a super seq of s2", s1, s2, expect, res)
	}

}

func isSubsequence(s, x string) bool {
	var i int
	var j int

	for i < len(s) && j < len(x) {
		if s[i] != x[j] {
			i++
		} else {
			i++
			j++
		}
	}
	return i <= len(s) && j == len(x)
}

func TestSample1(t *testing.T) {
	str1 := "abac"
	str2 := "cab"
	expect := "cabac"
	runSample(t, str1, str2, expect)
}

func TestSample2(t *testing.T) {
	str1 := "bbbaaaba"
	str2 := "bbababbb"
	expect := "bbabaaababb"
	runSample(t, str1, str2, expect)
}
