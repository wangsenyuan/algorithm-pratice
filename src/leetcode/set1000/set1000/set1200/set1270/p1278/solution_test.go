package p1278

import "testing"

func runSample(t *testing.T, s string, k int, expect int) {
	res := palindromePartition(s, k)
	if res != expect {
		t.Errorf("Sample %s %d, expect %d, but got %d", s, k, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "abc", 2, 1)
}

func TestSample2(t *testing.T) {
	runSample(t, "aabbc", 3, 0)
}

func TestSample3(t *testing.T) {
	runSample(t, "leetcode", 8, 0)
}

func TestSample4(t *testing.T) {
	runSample(t, "oiwwhqjkb", 1, 4)
}
