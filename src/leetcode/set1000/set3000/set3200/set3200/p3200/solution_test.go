package p3200

import (
	"testing"
)

func runSample(t *testing.T, nums []int, k int, expect int) {
	res := maximumLength(nums, k)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	k := 2
	expect := 5
	runSample(t, nums, k, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{1, 4, 2, 3, 1, 4}
	k := 3
	expect := 4
	runSample(t, nums, k, expect)
}
