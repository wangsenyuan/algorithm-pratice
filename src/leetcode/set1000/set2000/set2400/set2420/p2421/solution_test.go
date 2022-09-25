package p2421

import "testing"

func runSample(t *testing.T, vals []int, edges [][]int, expect int) {
	res := numberOfGoodPaths(vals, edges)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	vals := []int{1, 3, 2, 1, 3}
	edges := [][]int{{0, 1}, {0, 2}, {2, 3}, {2, 4}}
	expect := 6
	runSample(t, vals, edges, expect)
}

func TestSample2(t *testing.T) {
	vals := []int{1, 1, 2, 2, 3}
	edges := [][]int{{0, 1}, {1, 2}, {2, 3}, {2, 4}}
	expect := 7
	runSample(t, vals, edges, expect)
}
