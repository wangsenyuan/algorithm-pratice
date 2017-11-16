package mbp

import "testing"

func TestSample1(t *testing.T) {
	matrix := [][]int{
		{0, 1, 1, 0, 0, 0},
		{1, 0, 0, 1, 0, 0},
		{0, 0, 1, 0, 0, 0},
		{0, 0, 1, 1, 0, 0},
		{0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 1},
	}

	res := FindPerfectMaximumMatch(matrix, 6, 6)

	if res != 5 {
		t.Errorf("sample should give 5, but is %d", res)
	}
}
