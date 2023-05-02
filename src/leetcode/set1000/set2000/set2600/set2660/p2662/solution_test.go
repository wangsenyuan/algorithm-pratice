package p2662

import "testing"

func runSample(t *testing.T, start, target []int, roads [][]int, expect int) {
	res := minimumCost(start, target, roads)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	start := []int{1, 1}
	target := []int{9, 3}
	roads := [][]int{
		{6, 3, 9, 1, 1}, {7, 1, 6, 3, 1}, {7, 3, 4, 2, 2}, {3, 3, 1, 1, 2},
	}
	expect := 10
	runSample(t, start, target, roads, expect)
}
