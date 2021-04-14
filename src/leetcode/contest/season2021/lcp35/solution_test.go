package lcp35

import "testing"

func runSample(t *testing.T, paths [][]int, cnt int, start int, end int, charge []int, expect int) {
	res := electricCarPlan(paths, cnt, start, end, charge)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	paths := [][]int{
		{1, 3, 3}, {3, 2, 1}, {2, 1, 3}, {0, 1, 4}, {3, 0, 5},
	}
	cnt := 6
	start := 1
	end := 0
	charge := []int{2, 10, 4, 1}
	expect := 43
	runSample(t, paths, cnt, start, end, charge, expect)
}

func TestSample2(t *testing.T) {
	paths := [][]int{
		{0, 4, 2}, {4, 3, 5}, {3, 0, 5}, {0, 1, 5}, {3, 2, 4}, {1, 2, 8},
	}
	cnt := 8
	start := 0
	end := 2
	charge := []int{4, 1, 1, 3, 2}
	expect := 38
	runSample(t, paths, cnt, start, end, charge, expect)
}
