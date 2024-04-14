package p3111

import "testing"

func runSample(t *testing.T, nums []int, expect int) {
	res := numberOfSubarrays(nums)

	if expect != int(res) {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{1, 4, 3, 3, 2}
	expect := 6
	runSample(t, nums, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{3, 3, 3}
	expect := 6
	runSample(t, nums, expect)
}
