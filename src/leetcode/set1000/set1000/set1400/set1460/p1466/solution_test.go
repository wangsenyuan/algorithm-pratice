package p1466

import "testing"

func runSample(t *testing.T, n int, conns [][]int, expect int) {
	res := minReorder(n, conns)

	if res != expect {
		t.Errorf("Sample %d %v, expect %d, but got %d", n, conns, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 6
	conns := [][]int{{0, 1}, {1, 3}, {2, 3}, {4, 0}, {4, 5}}
	expect := 3
	runSample(t, n, conns, expect)
}

func TestSample2(t *testing.T) {
	n := 5
	conns := [][]int{{1, 0}, {1, 2}, {3, 2}, {3, 4}}
	expect := 2
	runSample(t, n, conns, expect)
}
