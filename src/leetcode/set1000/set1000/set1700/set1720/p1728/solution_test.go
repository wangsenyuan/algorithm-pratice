package p1728

import "testing"

func runSample(t *testing.T, matrix [][]int, expect int) {
	res := largestSubmatrix(matrix)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	matrix := [][]int{{0, 0, 1}, {1, 1, 1}, {1, 0, 1}}
	expect := 4
	runSample(t, matrix, expect)
}

func TestSample2(t *testing.T) {
	matrix := [][]int{{0, 1, 0, 1, 1, 0}}
	expect := 3
	runSample(t, matrix, expect)
}

func TestSample3(t *testing.T) {
	matrix := [][]int{{1, 1, 0}, {1, 0, 1}}
	expect := 2
	runSample(t, matrix, expect)
}
