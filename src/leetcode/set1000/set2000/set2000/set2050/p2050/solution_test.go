package p2050

import "testing"

func runSample(t *testing.T, n int, relations [][]int, time []int, expect int) {
	res := minimumTime(n, relations, time)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	relations := [][]int{
		{1, 3}, {2, 3},
	}
	time := []int{3, 2, 5}
	expect := 8
	runSample(t, n, relations, time, expect)
}

func TestSample2(t *testing.T) {
	n := 5
	relations := [][]int{
		{1, 5}, {2, 5}, {3, 5}, {3, 4}, {4, 5},
	}
	time := []int{1, 2, 3, 4, 5}
	expect := 12
	runSample(t, n, relations, time, expect)
}
