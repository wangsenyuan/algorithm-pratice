package p2322

import "testing"

func runSample(t *testing.T, nums []int, edges [][]int, expect int) {
	res := minimumScore(nums, edges)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{1, 5, 5, 4, 11}
	edges := [][]int{
		{0, 1}, {1, 2}, {1, 3}, {3, 4},
	}
	expect := 9
	runSample(t, nums, edges, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{5, 5, 2, 4, 4, 2}
	edges := [][]int{
		{0, 1}, {1, 2}, {5, 2}, {4, 3}, {1, 3},
	}
	expect := 0
	runSample(t, nums, edges, expect)
}
