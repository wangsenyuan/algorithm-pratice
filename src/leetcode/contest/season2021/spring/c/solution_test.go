package c

import "testing"

func runSample(t *testing.T, nums []int, expect int) {
	res := magicTower(nums)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{100, 100, 100, -250, -60, -140, -50, -50, 100, 150}
	expect := 1
	runSample(t, nums, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{-200, -300, 400, 0}
	expect := -1
	runSample(t, nums, expect)
}

func TestSample3(t *testing.T) {
	nums := []int{-200, -200, 400, 0}
	expect := 2
	runSample(t, nums, expect)
}
