package p2195

import "testing"

func runSample(t *testing.T, nums []int, k int, expect int64) {
	res := minimalKSum(nums, k)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{1, 4, 25, 10, 25}
	k := 2
	var expect int64 = 5
	runSample(t, nums, k, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{5, 6}
	k := 6
	var expect int64 = 25
	runSample(t, nums, k, expect)
}
