package main

import "testing"

func runSample(t *testing.T, A []int, expect bool) {
	res := solve(A)

	if len(res) > 0 != expect {
		t.Fatalf("Sample expect %t, but got %v", expect, res)
	}

	if !expect {
		return
	}
	n := len(A)
	g := NewGraph(n, len(res)*2)

	for _, e := range res {
		u, v := e[0]-1, e[1]-1
		if A[u] == A[v] {
			t.Fatalf("Sample result %v, not correct", res)
		}
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	var dfs func(p int, u int) int
	dfs = func(p int, u int) int {
		sz := 1
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				sz += dfs(u, v)
			}
		}
		return sz
	}

	sz := dfs(0, 0)

	if sz != n {
		t.Fatalf("Sample result %v, not a tree", res)
	}
}

type Graph struct {
	nodes []int
	next  []int
	to    []int
	cur   int
}

func NewGraph(n int, e int) *Graph {
	nodes := make([]int, n)
	next := make([]int, e+1)
	to := make([]int, e+1)
	return &Graph{nodes, next, to, 0}
}

func (g *Graph) AddEdge(u, v int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
}

func TestSample1(t *testing.T) {
	A := []int{1, 2, 2, 1, 3}
	expect := true
	runSample(t, A, expect)
}

func TestSample2(t *testing.T) {
	A := []int{1, 1, 1}
	expect := false
	runSample(t, A, expect)
}

func TestSample3(t *testing.T) {
	A := []int{1, 1000, 101, 1000}
	expect := true
	runSample(t, A, expect)
}

func TestSample4(t *testing.T) {
	A := []int{1, 2, 3, 4}
	expect := true
	runSample(t, A, expect)
}
