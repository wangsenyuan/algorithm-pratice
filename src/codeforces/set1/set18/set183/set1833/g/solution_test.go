package main

import (
	"sort"
	"testing"
)

func runSample(t *testing.T, n int, edges [][]int, expect bool) {
	ok, res := solve(n, edges)

	if ok != expect {
		t.Fatalf("Sample expect %t, but got %t, %v", expect, ok, res)
	}

	if !expect {
		return
	}
	g := NewGraph(n, 2*n)

	sort.Ints(res)
	j := 0
	for i, cur := range edges {
		if j < len(res) && i == res[j]-1 {
			j++
			continue
		}
		g.AddEdge(cur[0]-1, cur[1]-1, i)
		g.AddEdge(cur[1]-1, cur[0]-1, i)
	}
	vis := make([]bool, n)
	var dfs func(u int) int
	dfs = func(u int) int {
		vis[u] = true
		sz := 1
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if !vis[v] {
				sz += dfs(v)
			}
		}
		return sz
	}

	for i := 0; i < n; i++ {
		if !vis[i] {
			sz := dfs(i)
			if sz != 3 {
				t.Fatalf("Sample result %v, not correct", res)
			}
		}
	}
}

func TestSample1(t *testing.T) {
	n := 9
	edges := [][]int{
		{1, 2},
		{4, 3},
		{7, 9},
		{5, 4},
		{4, 6},
		{3, 2},
		{8, 7},
		{1, 7},
	}
	expect := true
	runSample(t, n, edges, expect)
}

func TestSample2(t *testing.T) {
	n := 6
	edges := [][]int{
		{1, 2},
		{1, 3},
		{4, 3},
		{1, 5},
		{6, 1},
	}
	expect := false
	runSample(t, n, edges, expect)
}

func TestSample3(t *testing.T) {
	n := 9
	edges := [][]int{
		{2, 6},
		{6, 9},
		{9, 1},
		{9, 7},
		{1, 8},
		{7, 3},
		{8, 5},
		{4, 7},
	}
	expect := true
	runSample(t, n, edges, expect)
}
