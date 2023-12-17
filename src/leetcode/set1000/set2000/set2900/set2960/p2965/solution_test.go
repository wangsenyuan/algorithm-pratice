package p2965

import "testing"

func runSample(t *testing.T, nums []int, k int, expect int) {
	res := maxFrequencyScore(nums, int64(k))
	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{1, 2, 6, 4}
	k := 3
	expect := 3
	runSample(t, nums, k, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{1, 4, 4, 2, 4}
	k := 0
	expect := 3
	runSample(t, nums, k, expect)
}

func TestSample3(t *testing.T) {
	nums := []int{16, 2, 6, 20, 2, 18, 16, 8, 15, 19, 22, 29, 24, 2, 26, 19}
	k := 40
	expect := 11
	runSample(t, nums, k, expect)
}
