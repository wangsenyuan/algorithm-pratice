package p3070

import "testing"

func runSample(t *testing.T, nums []int, k int, edges [][]int, expect int) {
	res := maximumValueSum(nums, k, edges)

	if res != int64(expect) {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{1, 2, 1}
	k := 3
	edges := [][]int{
		{0, 1}, {0, 2},
	}
	expect := 6
	runSample(t, nums, k, edges, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{2, 3}
	k := 7
	edges := [][]int{
		{0, 1},
	}
	expect := 9
	runSample(t, nums, k, edges, expect)
}
