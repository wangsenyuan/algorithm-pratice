package p757

import "testing"

func runSample(t *testing.T, intervals [][]int, expect int) {
	res := intersectionSizeTwo(intervals)
	if res != expect {
		t.Errorf("sample: %v; should give %d, but got %d", intervals, expect, res)
	}
}

func TestSample1(t *testing.T) {
	intervals := [][]int{
		{1, 3}, {1, 4}, {2, 5}, {3, 5},
	}
	expect := 3
	runSample(t, intervals, expect)
}

func TestSample2(t *testing.T) {
	intervals := [][]int{
		{1, 2}, {2, 3}, {2, 4}, {4, 5},
	}
	expect := 5
	runSample(t, intervals, expect)
}
