package p2570

import "testing"

func runSample(t *testing.T, lcp [][]int, expect string) {
	res := findTheString(lcp)

	if res != expect {
		t.Errorf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	lcp := [][]int{
		{4, 0, 2, 0},
		{0, 3, 0, 1},
		{2, 0, 2, 0},
		{0, 1, 0, 1},
	}
	expect := "abab"
	runSample(t, lcp, expect)
}

func TestSample2(t *testing.T) {
	lcp := [][]int{
		{4, 3, 2, 1},
		{3, 3, 2, 1},
		{2, 2, 2, 1},
		{1, 1, 1, 1},
	}
	expect := "aaaa"
	runSample(t, lcp, expect)
}

func TestSample3(t *testing.T) {
	lcp := [][]int{
		{4, 3, 2, 1},
		{3, 3, 2, 1},
		{2, 2, 2, 1},
		{1, 1, 1, 1},
	}
	expect := "aaaa"
	runSample(t, lcp, expect)
}

func TestSample4(t *testing.T) {
	lcp := [][]int{
		{4, 3, 2, 1},
		{3, 3, 2, 1},
		{2, 2, 2, 1},
		{1, 1, 1, 3},
	}
	expect := ""
	runSample(t, lcp, expect)
}

func TestSample5(t *testing.T) {
	lcp := [][]int{
		{0},
	}
	expect := ""
	runSample(t, lcp, expect)
}
