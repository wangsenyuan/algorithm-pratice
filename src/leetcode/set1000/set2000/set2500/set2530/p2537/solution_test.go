package p2537

import "testing"

func runSample(t *testing.T, nums []int, k int, expect int) {
	res := countGood(nums, k)

	if int(res) != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{1, 1, 1, 1, 1}
	k := 10
	expect := 1
	runSample(t, nums, k, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{3, 1, 4, 3, 2, 2, 4}
	k := 2
	expect := 4
	runSample(t, nums, k, expect)
}

func TestSample3(t *testing.T) {
	nums := []int{2, 3, 1, 3, 2, 3, 3, 3, 1, 1, 3, 2, 2, 2}
	k := 18
	expect := 9
	runSample(t, nums, k, expect)
}
