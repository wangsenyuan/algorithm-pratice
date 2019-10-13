package p1224

import "testing"

func runSample(t *testing.T, nums []int, expect int) {
	res := maxEqualFreq(nums)
	if res != expect {
		t.Errorf("Sample %v, expect %d, but got %d", nums, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, []int{2, 2, 1, 1, 5, 3, 3, 5}, 7)
}

func TestSample2(t *testing.T) {
	runSample(t, []int{1, 1, 1, 2, 2, 2, 3, 3, 3, 4, 4, 4, 5}, 13)
}

func TestSample3(t *testing.T) {
	runSample(t, []int{1, 1, 1, 2, 2, 2}, 5)
}

func TestSample4(t *testing.T) {
	runSample(t, []int{10, 2, 8, 9, 3, 8, 1, 5, 2, 3, 7, 6}, 8)
}

func TestSample5(t *testing.T) {
	runSample(t, []int{1, 2}, 2)
}

func TestSample6(t *testing.T) {
	runSample(t, []int{1, 1}, 2)
}
