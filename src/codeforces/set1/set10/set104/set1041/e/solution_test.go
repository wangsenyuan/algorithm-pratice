package main

import (
	"reflect"
	"sort"
	"testing"
)

func runSample(t *testing.T, n int, edges [][]int, expect bool) {
	res := solve(n, edges)

	if len(res) > 0 != expect {
		t.Fatalf("Sample expect %t, but got %v", expect, res)
	}

	if !expect {
		return
	}
	g := NewGraph(n+1, len(res)*2)

	for _, e := range res {
		u, v := e[0], e[1]
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}
	var arr [][]int
	var dfs func(p int, u int) int
	dfs = func(p int, u int) int {
		tmp := u
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if v != p {
				vv := dfs(u, v)
				tmp = max(tmp, vv)
				arr = append(arr, []int{vv, n})
			}
		}

		return tmp
	}
	dfs(n, n)
	for _, edge := range edges {
		u, v := edge[0], edge[1]
		edge[0] = min(u, v)
		edge[1] = max(v, u)
	}
	sort.Slice(edges, func(i, j int) bool {
		return edges[i][0] < edges[j][0]
	})
	sort.Slice(arr, func(i, j int) bool {
		return arr[i][0] < arr[j][0]
	})

	if !reflect.DeepEqual(arr, edges) {
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

func TestSample1(t *testing.T) {
	n := 4
	edges := [][]int{
		{3, 4},
		{1, 4},
		{3, 4},
	}
	expect := true
	runSample(t, n, edges, expect)
}

func TestSample2(t *testing.T) {
	n := 3
	edges := [][]int{
		{1, 3},
		{1, 3},
	}
	expect := false
	runSample(t, n, edges, expect)
}

func TestSample3(t *testing.T) {
	n := 3
	edges := [][]int{
		{1, 2},
		{2, 3},
	}
	expect := false
	runSample(t, n, edges, expect)
}
