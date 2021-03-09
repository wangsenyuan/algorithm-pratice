package p1787

import "testing"

func runSample(t *testing.T, nums []int, k int, expect int) {
	res := minChanges(nums, k)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{3, 4, 5, 2, 1, 7, 3, 4, 7}
	k := 3
	expect := 3
	runSample(t, nums, k, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{1, 2, 4, 1, 2, 5, 1, 2, 6}
	k := 3
	expect := 3
	runSample(t, nums, k, expect)
}

func TestSample3(t *testing.T) {
	nums := []int{26, 19, 19, 28, 13, 14, 6, 25, 28, 19, 0, 15, 25, 11}
	k := 3
	expect := 11
	runSample(t, nums, k, expect)
}
