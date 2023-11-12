package p2931

import "testing"

func runSample(t *testing.T, nums []int, expect int) {
	res := maximumStrongPairXor(nums)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	expect := 7
	runSample(t, nums, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{10, 100}
	expect := 0
	runSample(t, nums, expect)
}

func TestSample3(t *testing.T) {
	nums := []int{500, 520, 2500, 3000}
	expect := 1020
	runSample(t, nums, expect)
}
