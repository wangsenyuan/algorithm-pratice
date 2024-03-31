package p3099

import "testing"

func runSample(t *testing.T, nums []int, expect int) {
	res := countAlternatingSubarrays(nums)

	if res != int64(expect) {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{0, 1, 1, 1}
	expect := 5
	runSample(t, nums, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{1, 0, 1, 0}
	expect := 10
	runSample(t, nums, expect)
}
