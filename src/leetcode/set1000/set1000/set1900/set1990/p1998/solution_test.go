package p1998

import "testing"

func runSample(t *testing.T, nums []int, expect bool) {
	res := gcdSort(nums)

	if res != expect {
		t.Errorf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{7, 21, 3}
	expect := true
	runSample(t, nums, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{5, 2, 6, 2}
	expect := false
	runSample(t, nums, expect)
}

func TestSample3(t *testing.T) {
	nums := []int{10, 5, 9, 3, 15}
	expect := true
	runSample(t, nums, expect)
}
