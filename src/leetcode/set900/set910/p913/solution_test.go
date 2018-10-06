package p913

import "testing"

func runSample(t *testing.T, graph [][]int, expect int) {
	res := catMouseGame(graph)
	if res != expect {
		t.Errorf("Sample %v, expect %d, but got %d", graph, expect, res)
	}
}

func TestSample1(t *testing.T) {
	graph := [][]int{
		{2, 5}, {3}, {0, 4, 5}, {1, 4, 5}, {2, 3}, {0, 2, 3},
	}
	expect := 0
	runSample(t, graph, expect)
}
