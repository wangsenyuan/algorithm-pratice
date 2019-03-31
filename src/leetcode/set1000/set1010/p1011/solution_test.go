package p1011

import "testing"

func runSample(t *testing.T, A [][]int, expect int) {
	res := numEnclaves(A)
	if res != expect {
		t.Errorf("Sample %v, expect %d, but got %d", A, expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := [][]int{
		{0, 0, 0, 0},
		{1, 0, 1, 0},
		{0, 1, 1, 0},
		{0, 0, 0, 0},
	}
	runSample(t, A, 3)
}

func TestSample2(t *testing.T) {
	A := [][]int{
		{0, 1, 1, 0},
		{0, 0, 1, 0},
		{0, 0, 1, 0},
		{0, 0, 0, 0},
	}
	runSample(t, A, 0)
}
