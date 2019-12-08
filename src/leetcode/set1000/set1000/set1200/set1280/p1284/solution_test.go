package p1284

import "testing"

func runSample(t *testing.T, mat [][]int, expect int) {
	res := minFlips(mat)

	if res != expect {
		t.Errorf("Sample %v, expect %d, but got %d", mat, expect, res)
	}
}

func TestSample1(t *testing.T) {
	mat := [][]int{{0, 0}, {0, 1}}
	runSample(t, mat, 3)
}

func TestSample2(t *testing.T) {
	mat := [][]int{{0}}
	runSample(t, mat, 0)
}

func TestSample3(t *testing.T) {
	mat := [][]int{{1, 1, 1}, {1, 0, 1}, {0, 0, 0}}
	runSample(t, mat, 6)
}

func TestSample4(t *testing.T) {
	mat := [][]int{{1, 0, 0}, {1, 0, 0}}
	runSample(t, mat, -1)
}
