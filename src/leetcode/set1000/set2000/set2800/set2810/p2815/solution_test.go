package p2815

import "testing"

func runSample(t *testing.T, nums []int, k int, expect int) {
	res := maximumScore(nums, k)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{8, 3, 9, 3, 8}
	k := 2
	expect := 81
	runSample(t, nums, k, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{19, 12, 14, 6, 10, 18}
	k := 3
	expect := 4788
	runSample(t, nums, k, expect)
}
