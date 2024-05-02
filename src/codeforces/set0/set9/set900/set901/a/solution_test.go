package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, a []int, perfect bool) {
	res, trs := solve(a)

	if res != perfect {
		t.Fatalf("Sample expect %t, but got %t", perfect, res)
	}
	if perfect {
		return
	}

	for _, tr := range trs {
		b := countLevels(tr)

		if !reflect.DeepEqual(a, b) {
			t.Fatalf("Sample result %v, not correct", tr)
		}
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
	e++
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

func countLevels(p []int) []int {
	n := len(p)
	tr := NewGraph(n, n)

	for i := 1; i < n; i++ {
		tr.AddEdge(p[i]-1, i)
	}
	res := make([]int, n)
	var dfs func(u int, d int) int

	dfs = func(u int, d int) int {
		res[d]++
		var h int
		for i := tr.nodes[u]; i > 0; i = tr.next[i] {
			v := tr.to[i]
			h = max(h, dfs(v, d+1))
		}
		h++
		return h
	}
	h := dfs(0, 0)

	return res[:h]
}

func TestSample1(t *testing.T) {
	a := []int{1, 1, 1}
	expect := true
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{1, 2, 2}
	expect := false
	runSample(t, a, expect)
}
