package p2754

import "testing"

func runSample(t *testing.T, nums []int, expect int) {
	res := sumImbalanceNumbers(nums)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{2, 3, 1, 4}
	expect := 3
	runSample(t, nums, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{1, 3, 3, 3, 5}
	expect := 8
	runSample(t, nums, expect)
}
