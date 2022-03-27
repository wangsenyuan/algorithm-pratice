package p2218

import "testing"

func runSample(t *testing.T, piles [][]int, k int, expect int) {
	res := maxValueOfCoins(piles, k)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	piles := [][]int{
		{1, 100, 3}, {7, 8, 9},
	}
	k := 2
	expect := 101
	runSample(t, piles, k, expect)
}

func TestSample2(t *testing.T) {
	piles := [][]int{
		{100}, {100}, {100}, {100}, {100}, {100}, {1, 1, 1, 1, 1, 1, 700},
	}
	k := 7
	expect := 706
	runSample(t, piles, k, expect)
}

func TestSample3(t *testing.T) {
	piles := [][]int{
		{37, 88}, {51, 64, 65, 20, 95, 30, 26}, {9, 62, 20}, {44},
	}
	k := 9
	expect := 494
	runSample(t, piles, k, expect)
}
