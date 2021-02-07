package p1755

import "testing"

func runSample(t *testing.T, nums []int, goal int, expect int) {
	res := minAbsDifference(nums, goal)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{5, -7, 3, 5}
	goal := 6
	expect := 0
	runSample(t, nums, goal, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{7, -9, 15, -2}
	goal := -5
	expect := 1
	runSample(t, nums, goal, expect)
}

func TestSample3(t *testing.T) {
	nums := []int{7, -9, 15, -2}
	goal := -5
	expect := 1
	runSample(t, nums, goal, expect)
}

func TestSample4(t *testing.T) {
	nums := []int{1, 2, 3}
	goal := -7
	expect := 7
	runSample(t, nums, goal, expect)
}

func TestSample5(t *testing.T) {
	nums := []int{-7933, -1642, -6137, 6234, 4728, 5474, 2439}
	goal := -428059487
	expect := 428043775
	runSample(t, nums, goal, expect)
}
