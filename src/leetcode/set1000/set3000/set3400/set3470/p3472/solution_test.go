package p3472

import "testing"

func runSample(t *testing.T, s string, k int, expect int) {
	res := longestPalindromicSubsequence(s, k)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "abced"
	k := 2
	expect := 3
	runSample(t, s, k, expect)
}

func TestSample2(t *testing.T) {
	s := "aaazzz"
	k := 4
	expect := 6
	runSample(t, s, k, expect)
}


func TestSample3(t *testing.T) {
	s := "a"
	k := 20
	expect := 1
	runSample(t, s, k, expect)
}
