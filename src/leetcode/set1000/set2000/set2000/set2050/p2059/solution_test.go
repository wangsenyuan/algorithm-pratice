package p2059

import "testing"

func runSample(t *testing.T, nums []int, start int, goal int, expect int) {
	res := minimumOperations(nums, start, goal)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{1, 3}
	start := 6
	goal := 4

	expect := 2
	runSample(t, nums, start, goal, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{2, 4, 12}
	start := 2
	goal := 12

	expect := 2
	runSample(t, nums, start, goal, expect)
}

func TestSample3(t *testing.T) {
	nums := []int{3, 5, 7}
	start := 0
	goal := -4

	expect := 2
	runSample(t, nums, start, goal, expect)
}

func TestSample4(t *testing.T) {
	nums := []int{2, 8, 16}
	start := 0
	goal := 1

	expect := -1
	runSample(t, nums, start, goal, expect)
}

func TestSample5(t *testing.T) {
	nums := []int{1}
	start := 0
	goal := 3

	expect := 3
	runSample(t, nums, start, goal, expect)
}
