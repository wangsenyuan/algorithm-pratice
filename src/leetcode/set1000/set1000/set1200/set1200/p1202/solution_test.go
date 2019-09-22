package p1202

import "testing"

func runSample(t *testing.T, s string, pairs [][]int, expect string) {
	res := smallestStringWithSwaps(s, pairs)
	if res != expect {
		t.Errorf("Sample %s %v, expect %s, but got %s", s, pairs, expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "dcab"
	pairs := [][]int{{0, 3}, {1, 2}}
	expect := "bacd"
	runSample(t, s, pairs, expect)
}

func TestSample2(t *testing.T) {
	s := "dcab"
	pairs := [][]int{{0, 3}, {1, 2}, {0, 2}}
	expect := "abcd"
	runSample(t, s, pairs, expect)
}

func TestSample3(t *testing.T) {
	s := "cba"
	pairs := [][]int{{1, 2}, {0, 1}}
	expect := "abc"
	runSample(t, s, pairs, expect)
}
