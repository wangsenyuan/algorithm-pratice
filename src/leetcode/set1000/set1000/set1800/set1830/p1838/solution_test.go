package p1838

import "testing"

func runSample(t *testing.T, nums []int, k int, expect int) {
	res := maxFrequency(nums, k)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{1, 2, 4}
	k := 5
	expect := 3
	runSample(t, nums, k, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{1, 4, 8, 13}
	k := 5
	expect := 2
	runSample(t, nums, k, expect)
}

func TestSample3(t *testing.T) {
	nums := []int{3, 9, 6}
	k := 2
	expect := 1
	runSample(t, nums, k, expect)
}
