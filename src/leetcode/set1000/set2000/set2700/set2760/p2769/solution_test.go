package p2769

import "testing"

func runSample(t *testing.T, nums []int, k int, expect bool) {
	res := checkArray(nums, k)

	if res != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{2, 2, 3, 1, 1, 0}
	k := 3
	expect := true
	runSample(t, nums, k, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{1, 3, 1, 1}
	k := 2
	expect := false
	runSample(t, nums, k, expect)
}
