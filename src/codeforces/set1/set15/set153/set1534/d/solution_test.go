package main

import (
	"testing"
)

func runSample(t *testing.T, n int, edges [][]int) {
	g := NewGraph(n, 2*n)
	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	d := make([]int, n)

	var dfs func(p int, u int)

	dfs = func(p int, u int) {
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				d[v] = d[u] + 1
				dfs(u, v)
			}
		}
	}
	var cnt int

	ask := func(r int) []int {
		cnt++
		if cnt > (n+1)/2 {
			panic("asked too much")
		}
		r--
		d[r] = 0
		dfs(-1, r)
		return d
	}

	res := solve(n, ask)

	g2 := make([]map[int]bool, n)
	for i := 0; i < n; i++ {
		g2[i] = make(map[int]bool)
	}

	for _, e := range res {
		u, v := e[0]-1, e[1]-1
		g2[u][v] = true
		g2[v][u] = true
	}

	var validate func(p int, u int) bool
	validate = func(p int, u int) bool {
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if !g2[u][v] {
				return false
			}
			if p != v {
				if !validate(u, v) {
					return false
				}
			}
		}
		return true
	}

	if !validate(-1, 0) {
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
	g := new(Graph)
	g.nodes = make([]int, n)
	g.next = make([]int, e+2)
	g.to = make([]int, e+2)
	g.cur = 0
	return g
}

func (g *Graph) AddEdge(u, v int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
}

func TestSample1(t *testing.T) {
	n := 4
	edges := [][]int{
		{1, 2},
		{2, 3},
		{2, 4},
	}
	runSample(t, n, edges)
}

func TestSample2(t *testing.T) {
	n := 5
	edges := [][]int{
		{1, 3},
		{3, 5},
		{5, 4},
		{4, 2},
	}
	runSample(t, n, edges)
}

func TestSample3(t *testing.T) {
	n := 7
	edges := [][]int{
		{1, 2},
		{2, 3},
		{2, 4},
		{4, 5},
		{4, 6},
		{3, 7},
	}
	runSample(t, n, edges)
}

func TestSample4(t *testing.T) {
	n := 10
	edges := [][]int{
		{1, 2},
		{2, 3},
		{2, 4},
		{4, 5},
		{4, 6},
		{3, 7},
		{6, 8},
		{8, 9},
		{3, 10},
	}
	runSample(t, n, edges)
}
