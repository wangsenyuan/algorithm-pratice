package a

import "testing"

func runSample(t *testing.T, nums []int, target, expect int) {
	res := purchasePlans(nums, target)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{2, 3, 5, 5}
	target := 6
	expect := 1
	runSample(t, nums, target, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{2, 2, 1, 9}
	target := 10
	expect := 4
	runSample(t, nums, target, expect)
}
