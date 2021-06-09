package p1886

import "testing"

func runSample(t *testing.T, mat [][]int, target [][]int, expect bool) {
	res := findRotation(mat, target)

	if res != expect {
		t.Errorf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	mat := [][]int{
		{0, 1}, {1, 0},
	}

	target := [][]int{
		{1, 0}, {0, 1},
	}
	runSample(t, mat, target, true)
}

func TestSample2(t *testing.T) {
	mat := [][]int{
		{0, 1}, {1, 1},
	}

	target := [][]int{
		{1, 0}, {0, 1},
	}
	runSample(t, mat, target, false)
}

func TestSample3(t *testing.T) {
	mat := [][]int{
		{0, 0, 0}, {0, 1, 0}, {1, 1, 1},
	}

	target := [][]int{
		{1, 1, 1}, {0, 1, 0}, {0, 0, 0},
	}
	runSample(t, mat, target, true)
}
