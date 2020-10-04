package p1579

import "testing"

func runSample(t *testing.T, n int, edges [][]int, expect int) {
	res := maxNumEdgesToRemove(n, edges)

	if res != expect {
		t.Errorf("Sample %d %v, expect %d, but got %d", n, edges, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 4
	edges := [][]int{
		{3, 1, 2}, {3, 2, 3}, {1, 1, 3}, {1, 2, 4}, {1, 1, 2}, {2, 3, 4},
	}
	expect := 2
	runSample(t, n, edges, expect)
}

func TestSample2(t *testing.T) {
	n := 4
	edges := [][]int{
		{3, 1, 2}, {3, 2, 3}, {1, 1, 4}, {2, 1, 4},
	}
	expect := 0
	runSample(t, n, edges, expect)
}

func TestSample3(t *testing.T) {
	n := 4
	edges := [][]int{
		{3, 2, 3}, {1, 1, 2}, {2, 3, 4},
	}
	expect := -1
	runSample(t, n, edges, expect)
}
