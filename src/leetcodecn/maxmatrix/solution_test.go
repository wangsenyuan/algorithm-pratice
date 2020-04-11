package maxmatrix

import "testing"

func runSample(t *testing.T, matrix [][]int, expect []int) {
	res := getMaxMatrix(matrix)

	var sum int
	for i := expect[0]; i <= expect[2]; i++ {
		for j := expect[1]; j <= expect[3]; j++ {
			sum += matrix[i][j]
		}
	}
	var sum2 int

	for i := res[0]; i <= res[2]; i++ {
		for j := res[1]; j <= res[3]; j++ {
			sum2 += matrix[i][j]
		}
	}

	if sum != sum2 {
		t.Errorf("Sample %v, expect %v, but got %v", matrix, expect, res)
	}
}

func TestSample1(t *testing.T) {
	matrix := [][]int{
		{-1, 0},
		{0, -1},
	}
	expect := []int{0, 1, 0, 1}
	runSample(t, matrix, expect)
}

func TestSample2(t *testing.T) {
	matrix := [][]int{
		{-4, -5},
	}
	expect := []int{0, 0, 0, 0}
	runSample(t, matrix, expect)
}
