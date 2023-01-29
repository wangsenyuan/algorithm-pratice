package p2550

import "testing"

func runSample(t *testing.T, nums []int, expect int) {
	res := int(countQuadruplets(nums))

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{1, 3, 2, 4, 5}
	expect := 2
	runSample(t, nums, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{1, 2, 3, 4}
	expect := 0
	runSample(t, nums, expect)
}
