package p1981

import "testing"

func runSample(t *testing.T, mat [][]int, target int, expect int) {
	res := minimizeTheDifference(mat, target)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	mat := [][]int{
		{1, 2, 3}, {4, 5, 6}, {7, 8, 9},
	}
	target := 13
	expect := 0
	runSample(t, mat, target, expect)
}

func TestSample2(t *testing.T) {
	mat := [][]int{
		{1}, {2}, {3},
	}
	target := 100
	expect := 94
	runSample(t, mat, target, expect)
}

func TestSample3(t *testing.T) {
	mat := [][]int{
		{1, 2, 9, 8, 7},
	}
	target := 6
	expect := 1
	runSample(t, mat, target, expect)
}
