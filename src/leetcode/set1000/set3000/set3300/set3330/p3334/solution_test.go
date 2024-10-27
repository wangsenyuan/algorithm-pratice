package p3334

import "testing"

func runSample(t *testing.T, nums []int, expect int) {
	res := maxScore(nums)

	if res != int64(expect) {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{2, 4, 8, 16}
	expect := 64
	runSample(t, nums, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	expect := 60
	runSample(t, nums, expect)
}
