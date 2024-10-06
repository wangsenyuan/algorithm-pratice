package p3309

import "testing"

func runSample(t *testing.T, nums []int, expect int) {
	res := maxGoodNumber(nums)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{1, 2, 3}
	expect := 30
	runSample(t, nums, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{2, 8, 16}
	expect := 1296
	runSample(t, nums, expect)
}

func TestSample3(t *testing.T) {
	nums := []int{1, 11, 5}
	expect := 221
	runSample(t, nums, expect)
}
