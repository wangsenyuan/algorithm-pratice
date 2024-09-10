package main

import "testing"

func runSample(t *testing.T, n int, d int, expect bool) {
	res := solve(n, d)

	if len(res) > 0 != expect {
		t.Fatalf("Sample expect %t, but got %v", expect, res)
	}

	if !expect {
		return
	}

	g := NewGraph(n, n)

	for i := 0; i < n-1; i++ {
		if res[i] <= 0 || res[i] > n {
			t.Fatalf("Sample result %v, not correct", res)
		}
		g.AddEdge(res[i]-1, i+1)
	}

	dep := make([]int, n)

	var dfs func(u int)
	dfs = func(u int) {
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			dep[v] = dep[u] + 1
			dfs(v)
		}
	}

	dfs(0)
	var sum int
	for i := 1; i < n; i++ {
		if dep[i] == 0 {
			t.Fatalf("Sample result %v, not a tree", res)
		}
		sum += dep[i]
	}

	if sum != d {
		t.Fatalf("Sample result %v, not correct", res)
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
	next := make([]int, e)
	to := make([]int, e)
	return &Graph{nodes, next, to, 0}
}

func (g *Graph) AddEdge(u, v int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
}

func TestSample1(t *testing.T) {
	n, d := 5, 7
	expect := true
	runSample(t, n, d, expect)
}

func TestSample2(t *testing.T) {
	n, d := 10, 19
	expect := true
	runSample(t, n, d, expect)
}

func TestSample3(t *testing.T) {
	n, d := 10, 18
	expect := false
	runSample(t, n, d, expect)
}
