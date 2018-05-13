package p835

import "testing"

func runSample1(t *testing.T, A, B [][]int, expect int) {
	res := largestOverlap(A, B)
	if res != expect {
		t.Errorf("Sample %v %v, expect %d, but got %d", A, B, expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := [][]int{
		{1, 1, 0},
		{0, 1, 0},
		{0, 1, 0},
	}
	B := [][]int{
		{0, 0, 0},
		{0, 1, 1},
		{0, 0, 1},
	}

	expect := 3
	runSample1(t, A, B, expect)
}

func TestSample2(t *testing.T) {
	A := [][]int{
		{1, 0},
		{0, 0},
	}
	B := [][]int{
		{0, 1},
		{1, 0},
	}
	expect := 1
	runSample1(t, A, B, expect)
}

func TestSample3(t *testing.T) {
	A := [][]int{
		{1, 1, 1},
		{1, 0, 0},
		{0, 1, 1}}
	B := [][]int{
		{1, 1, 0},
		{1, 1, 1},
		{1, 1, 0}}
	expect := 4
	runSample1(t, A, B, expect)
}
