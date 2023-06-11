package p2730

import "testing"

func runSample(t *testing.T, nums []int, s string, d int, expect int) {
	res := sumDistance(nums, s, d)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{-2, 0, 2}
	s := "RLL"
	d := 3
	expect := 8
	runSample(t, nums, s, d, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{1, 0}
	s := "RL"
	d := 2
	expect := 5
	runSample(t, nums, s, d, expect)
}
