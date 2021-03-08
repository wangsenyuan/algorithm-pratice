package p1784

import "testing"

func runSample(t *testing.T, nums []int, limit int, goal int, expect int) {
	res := minElements(nums, limit, goal)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{1, -1, 1}
	limit := 3
	goal := -4
	expect := 2
	runSample(t, nums, limit, goal, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{1, -10, 9, 1}
	limit := 100
	goal := 0
	expect := 1
	runSample(t, nums, limit, goal, expect)
}
