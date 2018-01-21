package p768

import "testing"

func runSample(t *testing.T, nums []int, expect int) {
	res := maxChunksToSorted(nums)
	if res != expect {
		t.Errorf("sample %v, expect %d, but got %d", nums, expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{5, 4, 3, 2, 1}
	runSample(t, nums, 1)
}

func TestSample2(t *testing.T) {
	nums := []int{2, 1, 3, 4, 4}
	runSample(t, nums, 4)
}

func TestSample3(t *testing.T) {
	nums := []int{1, 1, 0, 0, 1}
	runSample(t, nums, 2)
}

func TestSample4(t *testing.T) {
	nums := []int{0, 0, 1, 1, 1}
	runSample(t, nums, 5)
}

func TestSample5(t *testing.T) {
	nums := []int{0, 3, 0, 3, 2}
	runSample(t, nums, 2)
}

func TestSample6(t *testing.T) {
	nums := []int{5, 1, 1, 8, 1, 6, 5, 9, 7, 8}
	runSample(t, nums, 1)
}
