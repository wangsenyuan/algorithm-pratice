package p886

import "testing"

func runSample(t *testing.T, edges [][]int, M int, N int, expect int) {
	res := reachableNodes(edges, M, N)

	if expect != res {
		t.Errorf("Sample %v %d %d, expect %d, but got %d", edges, M, N, expect, res)
	}
}

func TestSample1(t *testing.T) {
	edges := [][]int{{0, 1, 10}, {0, 2, 1}, {1, 2, 2}}
	M, N := 6, 3
	expect := 13
	runSample(t, edges, M, N, expect)
}

func TestSample2(t *testing.T) {
	edges := [][]int{{0, 1, 4}, {1, 2, 6}, {0, 2, 8}, {1, 3, 1}}
	M, N := 10, 4
	expect := 23
	runSample(t, edges, M, N, expect)
}
