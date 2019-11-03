package p1245

import "testing"

func runSample(t *testing.T, edges [][]int, expect int) {
	res := treeDiameter(edges)

	if res != expect {
		t.Errorf("Sample %v, expect %d, but got %d", edges, expect, res)
	}
}

func TestSample1(t *testing.T) {
	edges := [][]int{{0, 1}, {0, 2}}
	expect := 2
	runSample(t, edges, expect)
}

func TestSample2(t *testing.T) {
	edges := [][]int{{0, 1}, {1, 2}, {2, 3}, {1, 4}, {4, 5}}
	expect := 4
	runSample(t, edges, expect)
}
