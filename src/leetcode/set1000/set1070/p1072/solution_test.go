package p1072

import "testing"

func runSample(t *testing.T, matrix [][]int, expect int) {
	res := maxEqualRowsAfterFlips(matrix)
	if res != expect {
		t.Errorf("Sample %v, expect %d, but got %d", matrix, expect, res)
	}
}

func TestSample1(t *testing.T) {
	matrix := [][]int{
		{0, 1}, {1, 0},
	}
	runSample(t, matrix, 2)
}

func TestSample2(t *testing.T) {
	matrix := [][]int{
		{0, 0, 0}, {0, 0, 1}, {1, 1, 0},
	}
	runSample(t, matrix, 2)
}
