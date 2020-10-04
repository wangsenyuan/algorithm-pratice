package p1590

import "testing"

func runSample(t *testing.T, p int, arr []int, expect int) {
	res := minSubarray(arr, p)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{3, 1, 4, 2}
	p := 6
	expect := 1
	runSample(t, p, nums, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{6, 3, 5, 2}
	p := 9
	expect := 2
	runSample(t, p, nums, expect)
}

func TestSample3(t *testing.T) {
	nums := []int{1, 2, 3}
	p := 7
	expect := -1
	runSample(t, p, nums, expect)
}
