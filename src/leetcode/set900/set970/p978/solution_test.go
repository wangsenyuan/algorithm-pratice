package p978

import "testing"

func runSample(t *testing.T, A []int, expect int) {
	res := maxTurbulenceSize(A)
	if res != expect {
		t.Errorf("Sample %v, expect %d, but got %d", A, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, []int{9, 4, 2, 10, 7, 8, 8, 1, 9}, 5)
}

func TestSample2(t *testing.T) {
	runSample(t, []int{4, 8, 12, 16}, 2)
}

func TestSample3(t *testing.T) {
	runSample(t, []int{100}, 1)
}
