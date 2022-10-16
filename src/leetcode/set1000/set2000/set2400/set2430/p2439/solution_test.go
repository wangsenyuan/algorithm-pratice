package p2439

import "testing"

func runSample(t *testing.T, nums []int, edges [][]int, expect int) {
	res := componentValue(nums, edges)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{6, 2, 2, 2, 6}
	edges := [][]int{{0, 1}, {1, 2}, {1, 3}, {3, 4}}
	expect := 2
	runSample(t, nums, edges, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{1, 2, 1, 1, 1}
	edges := [][]int{{0, 1}, {1, 3}, {3, 4}, {4, 2}}
	expect := 1
	runSample(t, nums, edges, expect)
}
