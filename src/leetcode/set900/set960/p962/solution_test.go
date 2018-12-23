package p962

import "testing"

func runSample(t *testing.T, A []int, expect int) {
	res := maxWidthRamp(A)

	if res != expect {
		t.Errorf("Sample %v, expect %d, but got %d", A, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, []int{6, 0, 8, 2, 1, 5}, 4)
}

func TestSample2(t *testing.T) {
	runSample(t, []int{9, 8, 1, 0, 1, 9, 4, 0, 4, 1}, 7)
}

func TestSample3(t *testing.T) {
	runSample(t, []int{10, 10, 10, 7, 1, 6, 2, 1, 7}, 5)
}
