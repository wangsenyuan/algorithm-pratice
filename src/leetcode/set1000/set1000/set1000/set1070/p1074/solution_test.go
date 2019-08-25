package p1074

import "testing"

func runSample(t *testing.T, matrix [][]int, target int, expect int) {
	res := numSubmatrixSumTarget(matrix, target)

	if res != expect {
		t.Errorf("Sample %v %d, expect %d, but got %d", matrix, target, expect, res)
	}
}

func TestSample1(t *testing.T) {
	matrix := [][]int{
		{0, 1, 0}, {1, 1, 1}, {0, 1, 0},
	}

	target := 0
	expect := 4
	runSample(t, matrix, target, expect)
}

func TestSample2(t *testing.T) {
	matrix := [][]int{
		{1, -1}, {-1, 1},
	}

	target := 0
	expect := 5
	runSample(t, matrix, target, expect)
}

func TestSample3(t *testing.T) {
	matrix := [][]int{
		{0, 1, 1, 1, 0, 1}, {0, 0, 0, 0, 0, 1}, {0, 0, 1, 0, 0, 1}, {1, 1, 0, 1, 1, 0}, {1, 0, 0, 1, 0, 0},
	}

	target := 0
	expect := 43
	runSample(t, matrix, target, expect)
}
