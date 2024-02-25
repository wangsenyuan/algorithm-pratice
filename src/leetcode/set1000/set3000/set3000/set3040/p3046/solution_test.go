package p3046

import "testing"

func runSample(t *testing.T, nums []int, inds []int, expect int) {
	res := earliestSecondToMarkIndices(nums, inds)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{3, 2, 3}
	inds := []int{1, 3, 2, 2, 2, 2, 3}
	expect := 6
	runSample(t, nums, inds, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{0, 0, 1, 2}
	inds := []int{1, 2, 1, 2, 1, 2, 1, 2}
	expect := 7
	runSample(t, nums, inds, expect)
}

func TestSample3(t *testing.T) {
	nums := []int{2, 2}
	inds := []int{2, 2, 1, 2}
	expect := 4
	runSample(t, nums, inds, expect)
}

func TestSample4(t *testing.T) {
	nums := []int{0, 2}
	inds := []int{1, 1, 2, 2, 1}
	expect := 4
	runSample(t, nums, inds, expect)
}

func TestSample5(t *testing.T) {
	nums := []int{4, 0, 1, 0}
	inds := []int{4, 4, 4, 4, 2, 2, 1, 2, 4, 4, 1, 4, 2}
	expect := 8
	runSample(t, nums, inds, expect)
}
