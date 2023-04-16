package p2642

import "testing"

func TestSample1(t *testing.T) {
	n := 4
	edges := [][]int{
		{0, 2, 5}, {0, 1, 2}, {1, 2, 1}, {3, 0, 3},
	}
	g := Constructor(n, edges)

	w := g.ShortestPath(3, 2)

	if w != 6 {
		t.Fatalf("Graph should have shortest path 6 from 3 to 2, but got %d", w)
	}
	w = g.ShortestPath(0, 3)
	if w != -1 {
		t.Fatalf("Graph should have shortest path -1 from 0 to 3, but got %d", w)
	}
	g.AddEdge([]int{1, 3, 4})
	w = g.ShortestPath(0, 3)
	if w != 6 {
		t.Fatalf("Graph should have shortest path 6 from 0 to 3, but got %d", w)
	}
}
