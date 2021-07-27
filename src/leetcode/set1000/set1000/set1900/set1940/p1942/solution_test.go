package p1942

import "testing"

func runSample(t *testing.T, times [][]int, target int, expect int) {
	res := smallestChair(times, target)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	times := [][]int{
		{1, 4}, {2, 3}, {4, 6},
	}
	target := 1
	expect := 1
	runSample(t, times, target, expect)
}

func TestSample2(t *testing.T) {
	times := [][]int{
		{3, 10}, {1, 5}, {2, 6},
	}
	target := 0
	expect := 2
	runSample(t, times, target, expect)
}

func TestSample3(t *testing.T) {
	times := [][]int{
		{4, 5}, {12, 13}, {5, 6}, {1, 2}, {8, 9}, {9, 10}, {6, 7}, {3, 4}, {7, 8}, {13, 14}, {15, 16}, {14, 15}, {10, 11}, {11, 12}, {2, 3}, {16, 17},
	}
	target := 15
	expect := 0
	runSample(t, times, target, expect)
}
