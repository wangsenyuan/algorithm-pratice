package p945

import "testing"

func runSample(t *testing.T, A []int, expect int) {
	res := minIncrementForUnique(A)
	if res != expect {
		t.Errorf("Sample %v, expect %d, but got %d", A, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, []int{1, 2, 2}, 1)
}

func TestSample2(t *testing.T) {
	runSample(t, []int{3, 2, 1, 2, 1, 7}, 6)
}

func TestSample3(t *testing.T) {
	runSample(t, []int{0, 2, 2}, 1)
}
