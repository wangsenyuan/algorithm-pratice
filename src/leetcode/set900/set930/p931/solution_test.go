package p931

import "testing"

func runSample(t *testing.T, A [][]int, expect int) {
	res := minFallingPathSum(A)

	if res != expect {
		t.Errorf("Sample %v, expect %d, but got %d", A, expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := [][]int{
		{1, 2, 3}, {4, 5, 6}, {7, 8, 9},
	}
	expect := 12
	runSample(t, A, expect)
}
