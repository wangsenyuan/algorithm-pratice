package main

import "testing"

func runSample(t *testing.T, n int, edges [][]int, expect bool) {
	res := solve(n, edges)

	if len(res) > 0 != expect {
		t.Fatalf("Sample expect %t, but got %v", expect, res)
	}

	if !expect {
		return
	}
	if len(edges) != len(res) {
		t.Fatalf("Sample result %v, not correct", res)
	}

	type pair struct {
		first  int
		second int
	}

	routes := make(map[pair]int)

	for _, e := range edges {
		u, v := min(e[1], e[2]), max(e[1], e[2])
		key := pair{u, v}
		if e[0] == 1 {
			routes[key] = e[1]
		} else {
			routes[key] = 0
		}
	}

	g := NewGraph(n+1, len(edges))
	deg := make([]int, n+1)
	for _, cur := range res {
		u, v := min(cur[0], cur[1]), max(cur[0], cur[1])
		key := pair{u, v}
		if x, ok := routes[key]; !ok {
			t.Fatalf("Sample %v, have a non-existing edge", res)
		} else if x > 0 && x != cur[0] {
			t.Fatalf("Sample result %v, change direction of edge %v", res, key)
		}

		g.AddEdge(cur[0], cur[1])
		deg[cur[1]]++
	}

	que := make([]int, n)
	var head, tail int
	for i := 1; i <= n; i++ {
		if deg[i] == 0 {
			que[head] = i
			head++
		}
	}

	for tail < head {
		u := que[tail]
		tail++
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			deg[v]--
			if deg[v] == 0 {
				que[head] = v
				head++
			}
		}
	}

	for i := 1; i <= n; i++ {
		if deg[i] != 0 {
			t.Fatalf("Sample result %v, not correct", res)
		}
	}
}

func TestSample1(t *testing.T) {
	n := 3
	edges := [][]int{
		{0, 1, 3},
	}
	expect := true
	runSample(t, n, edges, expect)
}

func TestSample2(t *testing.T) {
	n := 5
	edges := [][]int{
		{0, 2, 1},
		{1, 1, 5},
		{1, 5, 4},
		{0, 5, 2},
		{1, 3, 5},
	}
	expect := true
	runSample(t, n, edges, expect)
}

func TestSample3(t *testing.T) {
	n := 4
	edges := [][]int{
		{1, 1, 2},
		{0, 4, 3},
		{1, 3, 1},
		{0, 2, 3},
		{1, 2, 4},
	}
	expect := true
	runSample(t, n, edges, expect)
}

func TestSample4(t *testing.T) {
	n := 4
	edges := [][]int{
		{1, 4, 1},
		{1, 1, 3},
		{0, 1, 2},
		{1, 2, 4},
		{1, 3, 2},
	}
	expect := false
	runSample(t, n, edges, expect)
}
