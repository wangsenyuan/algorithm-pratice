package p2963

import "testing"

func runSample(t *testing.T, nums []int, expect int) {
	res := numberOfGoodPartitions(nums)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{1, 2, 3, 4}
	expect := 8
	runSample(t, nums, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{1, 1, 1, 1}
	expect := 1
	runSample(t, nums, expect)
}

func TestSample3(t *testing.T) {
	nums := []int{1, 2, 1, 3}
	expect := 2
	runSample(t, nums, expect)
}
