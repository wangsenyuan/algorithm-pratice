package p1614

import "testing"

func runSample(t *testing.T, row, col []int) {
	a, b := make([]int, len(row)), make([]int, len(col))
	copy(a, row)
	copy(b, col)
	mat := restoreMatrix(a, b)

	for i := 0; i < len(row); i++ {
		var tmp int
		for j := 0; j < len(col); j++ {
			tmp += mat[i][j]
		}

		if tmp != row[i] {
			t.Fatalf("Sample result %v, not correct", mat)
		}
	}

	for j := 0; j < len(col); j++ {
		var tmp int
		for i := 0; i < len(row); i++ {
			tmp += mat[i][j]
		}
		if tmp != col[j] {
			t.Fatalf("Sample result %v, not correct", mat)
		}
	}
}

func TestSample1(t *testing.T) {
	row := []int{3, 8}
	col := []int{4, 7}
	runSample(t, row, col)
}

func TestSample2(t *testing.T) {
	row := []int{5, 7, 10}
	col := []int{8, 6, 8}
	runSample(t, row, col)
}

func TestSample3(t *testing.T) {
	row := []int{14, 9}
	col := []int{6, 9, 8}
	runSample(t, row, col)
}
