package p2134

import "testing"

func runSample(t *testing.T, arr []int, expect int) {
	res := minSwaps(arr)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{0, 1, 0, 1, 1, 0, 0}
	expect := 1
	runSample(t, nums, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{0, 1, 1, 1, 0, 0, 1, 1, 0}
	expect := 2
	runSample(t, nums, expect)
}

func TestSample3(t *testing.T) {
	nums := []int{1, 1, 0, 0, 1}
	expect := 0
	runSample(t, nums, expect)
}
