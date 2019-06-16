package p766

import "testing"

func runSample(t *testing.T, matrix [][]int, expect bool) {
	res := isToeplitzMatrix(matrix)
	if res != expect {
		t.Errorf("sample %v, expect %t, but got %t", matrix, expect, res)
	}
}

func TestSample1(t *testing.T) {
	matrix := [][]int{
		{1, 2, 3, 4}, {5, 1, 2, 3}, {9, 5, 1, 2},
	}
	runSample(t, matrix, true)
}

func TestSample2(t *testing.T) {
	matrix := [][]int{
		{1, 2}, {2, 2},
	}
	runSample(t, matrix, false)
}
