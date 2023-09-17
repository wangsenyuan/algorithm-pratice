package p2861

import "testing"

func runSample(t *testing.T, nums []int, expect int) {
	res := countWays(nums)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{1, 1}
	expect := 2
	runSample(t, nums, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{6, 0, 3, 3, 6, 7, 2, 7}
	expect := 3
	runSample(t, nums, expect)
}
