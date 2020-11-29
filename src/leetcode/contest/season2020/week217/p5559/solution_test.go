package p5559

import "testing"

func runSample(t *testing.T, arr []int, expect int) {
	res := minimumMountainRemovals(arr)

	if res != expect {
		t.Errorf("Sample %v, expect %d, but got %d", arr, expect, res)
	}
}

func TestSample1(t *testing.T) {
	arr := []int{1, 3, 1}
	expect := 0
	runSample(t, arr, expect)
}

func TestSample2(t *testing.T) {
	arr := []int{2, 1, 1, 5, 6, 2, 3, 1}
	expect := 3
	runSample(t, arr, expect)
}

func TestSample3(t *testing.T) {
	arr := []int{4, 3, 2, 1, 1, 2, 3, 1}
	expect := 4
	runSample(t, arr, expect)
}

func TestSample4(t *testing.T) {
	arr := []int{1, 2, 3, 4, 4, 3, 2, 1}
	expect := 1
	runSample(t, arr, expect)
}
