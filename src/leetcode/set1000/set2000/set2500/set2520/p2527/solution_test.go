package p2527

import "testing"

func runSample(t *testing.T, nums []int, expect int) {
	res := xorBeauty(nums)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{1, 4}
	expect := 5
	runSample(t, nums, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{15, 45, 20, 2, 34, 35, 5, 44, 32, 30}
	expect := 34
	runSample(t, nums, expect)
}
