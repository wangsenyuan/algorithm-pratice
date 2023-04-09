package p2616

import "testing"

func runSample(t *testing.T, nums []int, p int, expect int) {
	res := minimizeMax(nums, p)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{10, 1, 2, 7, 1, 3}
	p := 2
	expect := 1
	runSample(t, nums, p, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{4, 2, 1, 2}
	p := 1
	expect := 0
	runSample(t, nums, p, expect)
}

func TestSample3(t *testing.T) {
	nums := []int{3, 11, 4, 3, 5, 7, 4, 4, 5, 5}
	p := 3
	expect := 0
	runSample(t, nums, p, expect)
}
