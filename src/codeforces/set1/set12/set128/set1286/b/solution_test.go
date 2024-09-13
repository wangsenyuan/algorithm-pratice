package main

import (
	"sort"
	"testing"
)

func runSample(t *testing.T, nodes [][]int, expect bool) {
	res := solve(nodes)

	if len(res) > 0 != expect {
		t.Fatalf("Sample expect %t, but got %v", expect, res)
	}

	if !expect {
		return
	}
	n := len(nodes)
	g := NewGraph(n, n)
	cnt := make([]int, n)
	var root int
	for i := range nodes {
		p := nodes[i][0]
		if p > 0 {
			g.AddEdge(p-1, i)
		} else {
			root = i
		}
		cnt[i] = nodes[i][1]
	}

	dp := make([][]int, n)
	var dfs func(u int)

	dfs = func(u int) {
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			dfs(v)
			dp[u] = append(dp[u], dp[v]...)
		}
		dp[u] = append(dp[u], res[u])
		sort.Ints(dp[u])
		if dp[u][cnt[u]] != res[u] {
			t.Fatalf("Sample result %v, not correct", res)
		}
	}

	dfs(root)
}

func TestSample1(t *testing.T) {
	nodes := [][]int{
		{2, 0},
		{0, 2},
		{2, 0},
	}
	expect := true
	runSample(t, nodes, expect)
}

func TestSample2(t *testing.T) {
	nodes := [][]int{
		{0, 1},
		{1, 3},
		{2, 1},
		{3, 0},
		{2, 0},
	}
	expect := true
	runSample(t, nodes, expect)
}
