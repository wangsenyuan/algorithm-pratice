package p3174

import "testing"

func runSample(t *testing.T, nums []int, k int, expect int) {
	res := maximumLength(nums, k)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{1, 2, 1, 1, 3}
	k := 2
	expect := 4
	runSample(t, nums, k, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 1}
	k := 0
	expect := 2
	runSample(t, nums, k, expect)
}
