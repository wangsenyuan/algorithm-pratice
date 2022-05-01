package p2261

import "testing"

func runSample(t *testing.T, nums []int, k int, p int, expect int) {
	res := countDistinct(nums, k, p)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{2, 3, 3, 2, 2}
	k := 2
	p := 2
	expect := 11
	runSample(t, nums, k, p, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{99, 18, 2, 43, 91, 96, 39, 15, 47, 2, 141, 167, 133, 116, 27, 165, 177, 65, 33, 4, 113, 3, 154, 121, 163}
	k := 5
	p := 39
	expect := 324
	runSample(t, nums, k, p, expect)
}

func TestSample3(t *testing.T) {
	nums := []int{1, 9, 8, 7, 19}
	k := 1
	p := 6
	expect := 15
	runSample(t, nums, k, p, expect)
}
