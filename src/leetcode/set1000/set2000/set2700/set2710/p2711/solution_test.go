package p2711

import "testing"

func runSample(t *testing.T, mat [][]int, expect int) {
	res := maxIncreasingCells(mat)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	mat := [][]int{
		{3, 1}, {3, 4},
	}
	expect := 2
	runSample(t, mat, expect)
}

func TestSample2(t *testing.T) {
	mat := [][]int{
		{1, 1}, {1, 1},
	}
	expect := 1
	runSample(t, mat, expect)
}

func TestSample3(t *testing.T) {
	mat := [][]int{
		{3, 1, 6}, {-9, 5, 7},
	}
	expect := 4
	runSample(t, mat, expect)
}
