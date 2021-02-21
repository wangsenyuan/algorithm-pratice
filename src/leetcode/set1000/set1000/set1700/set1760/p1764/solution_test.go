package p1764

import "testing"

func runSample(t *testing.T, groups [][]int, nums []int, expect bool) {
	res := canChoose(groups, nums)

	if res != expect {
		t.Errorf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	groups := [][]int{
		{1, -1, -1}, {3, -2, 0},
	}
	nums := []int{1, -1, 0, 1, -1, -1, 3, -2, 0}
	expect := true
	runSample(t, groups, nums, expect)
}

func TestSample2(t *testing.T) {
	groups := [][]int{
		{10, -2}, {1, 2, 3, 4},
	}
	nums := []int{1, 2, 3, 4, 10, -2}
	expect := false
	runSample(t, groups, nums, expect)
}

func TestSample3(t *testing.T) {
	groups := [][]int{
		{1, 2, 3}, {3, 4},
	}
	nums := []int{7, 7, 1, 2, 3, 4, 7, 7}
	expect := false
	runSample(t, groups, nums, expect)
}
