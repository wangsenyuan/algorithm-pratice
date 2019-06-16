package p861

import "testing"

func runSample(t *testing.T, A [][]int, expect int) {
	res := matrixScore(A)

	if res != expect {
		t.Errorf("Sample %v, expect %d, but got %d", A, expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := [][]int{
		{0, 0, 1, 1}, {1, 0, 1, 0}, {1, 1, 0, 0},
	}
	expect := 39
	runSample(t, A, expect)
}
