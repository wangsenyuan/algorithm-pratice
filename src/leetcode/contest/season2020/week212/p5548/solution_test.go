package p5548

import "testing"

func runSample(t *testing.T, h [][]int, expect int) {
	res := minimumEffortPath(h)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	h := [][]int{{1, 2, 2}, {3, 8, 2}, {5, 3, 5}}
	expect := 2
	runSample(t, h, expect)
}

func TestSample2(t *testing.T) {
	h := [][]int{{1, 2, 3}, {3, 8, 4}, {5, 3, 5}}
	expect := 1
	runSample(t, h, expect)
}

func TestSample3(t *testing.T) {
	h := [][]int{{1, 2, 1, 1, 1}, {1, 2, 1, 2, 1}, {1, 2, 1, 2, 1}, {1, 2, 1, 2, 1}, {1, 1, 1, 2, 1}}
	expect := 0
	runSample(t, h, expect)
}

func TestSample4(t *testing.T) {
	h := [][]int{{1, 10, 6, 7, 9, 10, 4, 9}}
	expect := 9
	runSample(t, h, expect)
}
