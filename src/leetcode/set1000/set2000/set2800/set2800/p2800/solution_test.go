package p2800

import "testing"

func runSample(t *testing.T, nums []int, expect int) {
	res := countCompleteSubarrays(nums)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{1, 3, 1, 2, 2}
	expect := 4
	runSample(t, nums, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{5, 5, 5, 5}
	expect := 10
	runSample(t, nums, expect)
}
