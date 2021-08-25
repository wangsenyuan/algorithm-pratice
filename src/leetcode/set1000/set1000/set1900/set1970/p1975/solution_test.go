package p1975

import "testing"

func runSample(t *testing.T, mat [][]int, expect int64) {
	res := maxMatrixSum(mat)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	mat := [][]int{
		{1, -1}, {-1, 1},
	}
	var expect int64 = 4
	runSample(t, mat, expect)
}

func TestSample2(t *testing.T) {
	mat := [][]int{
		{1, 2, 3}, {-1, -2, -3}, {1, 2, 3},
	}
	var expect int64 = 16
	runSample(t, mat, expect)
}

func TestSample3(t *testing.T) {
	mat := [][]int{
		{-10000, -10000, -10000}, {-10000, -10000, -10000}, {-10000, -10000, -10000},
	}
	var expect int64 = 70000
	runSample(t, mat, expect)
}
