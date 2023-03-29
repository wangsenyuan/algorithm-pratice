package main

import "testing"

func runSample(t *testing.T, n int, E [][]int) {
	x, res := solve(n, E)

	if x == 1 {
		check1(t, n, E, res)
	} else {
		check2(t, n, E, res)
	}
}

func check1(t *testing.T, n int, E [][]int, color []int) {
	// 两条边不能有相同的颜色
	for _, e := range E {
		u, v := e[0], e[1]
		u--
		v--
		if color[u] == color[v] || color[u] < 1 || color[u] > 4 || color[v] < 1 || color[v] > 4 {
			t.Fatalf("Sample result %v, not correct, as edge %d - %d has same colo %d", color, u, v, color[u])
		}
	}
}

func check2(t *testing.T, n int, E [][]int, vex []int) {

	if len(vex)%2 != 1 {
		t.Fatalf("Sample result %v, not correct", vex)
	}

	g := NewGraph(n, 2*len(E))

	for _, e := range E {
		u, v := e[0]-1, e[1]-1
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	for j := 0; j < len(vex); j++ {
		u := vex[j] - 1
		v := vex[(j+1)%len(vex)] - 1
		ok := false
		for i := g.node[u]; i > 0 && !ok; i = g.next[i] {
			w := g.to[i]
			ok = v == w
		}
		if !ok {
			t.Fatalf("Sample result %v, no edge found between %d %d", vex, u, v)
		}
	}
}

func TestSample1(t *testing.T) {
	n := 6
	E := [][]int{
		{1, 2},
		{1, 3},
		{2, 3},
		{2, 4},
		{5, 2},
		{3, 5},
		{6, 5},
		{3, 6},
		{4, 5},
	}
	runSample(t, n, E)
}
