package problem4

import (
	"testing"
)

func runSample(t *testing.T, n int, m int, broken [][]int, expect int) {
	res := domino(n, m, broken)

	if res != expect {
		t.Errorf("Sample %d %d %v, expect %d, but got %d", n, m, broken, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 2
	m := 3
	broken := [][]int{{1, 0}, {1, 1}}
	expect := 2
	runSample(t, n, m, broken, expect)
}