package p2812

import "testing"

func runSample(t *testing.T, nums []int, m int, expect bool) {
	res := canSplitArray(nums, m)

	if res != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{2, 3, 3, 2, 3}
	m := 6
	expect := true
	runSample(t, nums, m, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{2, 1, 1, 3}
	m := 4
	expect := true
	runSample(t, nums, m, expect)
}
