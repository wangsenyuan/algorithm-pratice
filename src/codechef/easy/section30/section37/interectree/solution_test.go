package main

import "testing"

func runSample(t *testing.T, n int, edges [][]int, marked int) {
	ask := func(arr []int, k int) int {
		g := NewGraph(n, 2*n)

		rm := make(map[int]bool)

		for i := 0; i < k; i++ {
			rm[arr[i]-1] = true
		}

		for i, edge := range edges {
			if !rm[i] {
				u, v := edge[0], edge[1]
				u--
				v--
				g.AddEdge(u, v, i)
				g.AddEdge(v, u, i)
			}
		}

		return g.height_from(marked-1) - 1
	}

	res := solve(n, edges, ask)

	if res != marked {
		t.Errorf("Sample expect %d, but got %d", marked, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	edges := [][]int{{1, 2}, {2, 3}}
	marked := 2
	runSample(t, n, edges, marked)
}

func TestSample2(t *testing.T) {
	n := 4
	edges := [][]int{{1, 2}, {2, 3}, {3, 4}}
	marked := 3
	runSample(t, n, edges, marked)
}

func TestSample3(t *testing.T) {
	n := 4
	edges := [][]int{{1, 2}, {1, 3}, {1, 4}}
	marked := 2
	runSample(t, n, edges, marked)
}

func TestSample4(t *testing.T) {
	n := 6
	edges := [][]int{{1, 2}, {1, 3}, {4, 5}, {4, 6}, {1, 4}}
	marked := 3
	runSample(t, n, edges, marked)
}
