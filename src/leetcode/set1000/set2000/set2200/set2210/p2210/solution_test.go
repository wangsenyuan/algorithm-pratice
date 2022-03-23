package p2210

import "testing"

func runSample(t *testing.T, nums []int, expect int) {
	res := countHillValley(nums)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{2, 4, 1, 1, 6, 5}
	expect := 3
	runSample(t, nums, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{6, 6, 5, 5, 4, 1}
	expect := 0
	runSample(t, nums, expect)
}
