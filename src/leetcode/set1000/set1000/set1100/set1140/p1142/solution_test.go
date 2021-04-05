package p1142

import "testing"

func runSample(t *testing.T, text1, text2 string, expect int) {
	res := longestCommonSubsequence(text1, text2)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	text1 := "abcde"
	text2 := "ace"
	expect := 3
	runSample(t, text1, text2, expect)
}
