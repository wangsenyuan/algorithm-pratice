package p081

import "testing"

func runSample(t *testing.T, nums []int, target int, expect bool) {
	res := search(nums, target)

	if res != expect {
		t.Errorf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{2, 5, 6, 0, 0, 1, 2}
	target := 0
	expect := true
	runSample(t, nums, target, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{1, 0, 1, 1, 1}
	target := 0
	expect := true
	runSample(t, nums, target, expect)
}

func TestSample3(t *testing.T) {
	nums := []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 1, 1, 1, 1}
	target := 2
	expect := true
	runSample(t, nums, target, expect)
}
