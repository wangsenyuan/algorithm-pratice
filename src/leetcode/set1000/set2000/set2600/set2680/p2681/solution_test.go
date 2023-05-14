package p2681

import "testing"

func runSample(t *testing.T, nums []int, expect int) {
	res := sumOfPower(nums)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{2, 1, 4}
	expect := 141
	runSample(t, nums, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{1, 1, 1}
	expect := 7
	runSample(t, nums, expect)
}
