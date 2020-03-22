package p1388

import "testing"

func runSample(t *testing.T, slices []int, expect int) {
	res := maxSizeSlices(slices)

	if res != expect {
		t.Errorf("Sample %v, expect %d, but got %d", slices, expect, res)
	}
}

func TestSample1(t *testing.T) {
	slices := []int{1, 2, 3, 4, 5, 6}
	expect := 10
	runSample(t, slices, expect)
}

func TestSample2(t *testing.T) {
	slices := []int{4, 1, 2, 5, 8, 3, 1, 9, 7}
	expect := 21
	runSample(t, slices, expect)
}

func TestSample3(t *testing.T) {
	slices := []int{3, 2, 1}
	expect := 3
	runSample(t, slices, expect)
}
