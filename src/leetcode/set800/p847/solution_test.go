package p847

import "testing"

func runSample(t *testing.T, graph [][]int, expect int) {
	res := shortestPathLength(graph)
	if res != expect {
		t.Errorf("Sample %v, expect %d, but got %d", graph, expect, res)
	}
}

func TestSample1(t *testing.T) {
	graph := [][]int{
		{1, 2, 3}, {0}, {0}, {0},
	}
	expect := 4
	runSample(t, graph, expect)
}

func TestSample2(t *testing.T) {
	graph := [][]int{
		{1}, {0, 2, 4}, {1, 3, 4}, {2}, {1, 2},
	}
	expect := 4
	runSample(t, graph, expect)
}
