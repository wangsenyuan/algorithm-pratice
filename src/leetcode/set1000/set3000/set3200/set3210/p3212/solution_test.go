package p3212

import "testing"

func runSample(t *testing.T, grid [][]byte, expect int) {
	res := numberOfSubmatrices(grid)

	if expect != res {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	grid := [][]byte{
		{'X', 'Y', '.'}, {'Y', '.', '.'},
	}
	expect := 3
	runSample(t, grid, expect)
}
