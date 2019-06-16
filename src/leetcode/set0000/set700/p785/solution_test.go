package p785

import "testing"

func runSample(t *testing.T, graph [][]int, expect bool) {
	res := isBipartite(graph)

	if res != expect {
		t.Errorf("sample %v, expect %t, but got %t", graph, expect, res)
	}
}

func TestSample1(t *testing.T) {
	graph := [][]int{
		{1, 3}, {0, 2}, {1, 3}, {0, 2},
	}
	runSample(t, graph, true)
}

func TestSample2(t *testing.T) {
	graph := [][]int{
		{1, 2, 3}, {0, 2}, {1, 3}, {0, 2},
	}
	runSample(t, graph, false)
}
