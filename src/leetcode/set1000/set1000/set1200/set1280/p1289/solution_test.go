package p1289

import "testing"

func runSample(t *testing.T, arr [][]int, expect int) {
	res := minFallingPathSum(arr)
	if res != expect {
		t.Errorf("Sample %v, expect %d, but got %d", arr, expect, res)
	}
}

func TestSample1(t *testing.T) {
	arr := [][]int{
		{1, 2, 3}, {4, 5, 6}, {7, 8, 9},
	}

	expect := 13

	runSample(t, arr, expect)
}
