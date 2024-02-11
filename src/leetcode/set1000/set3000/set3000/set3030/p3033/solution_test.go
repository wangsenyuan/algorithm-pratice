package p3033

import "testing"

func runSample(t *testing.T, nums []int, pat []int, expect int) {
	res := countMatchingSubarrays(nums, pat)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6}
	pat := []int{1, 1}
	expect := 4
	runSample(t, nums, pat, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{1, 4, 4, 1, 3, 5, 5, 3}
	pat := []int{1, 0, -1}
	expect := 2
	runSample(t, nums, pat, expect)
}
