package main

import "testing"

func runSample(t *testing.T, k int, P []int) {
	res := solve(k, P)

	n := len(P) + 1
	g := NewGraph(n, n)

	for i := 0; i < len(P); i++ {
		p := P[i] - 1
		g.AddEdge(p, i+1)
	}
	X := make([]int, n)
	var dfs func(u int) int

	dfs = func(u int) int {
		X[u] = int(res[u] - '0')
		cur := X[u]
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			cur ^= dfs(v)
		}
		k -= cur
		return cur
	}
	dfs(0)

	if k != 0 {
		t.Fatalf("Sample result %v, not correct", res)
	}
}

func TestSample1(t *testing.T) {
	k := 0
	P := []int{1}
	runSample(t, k, P)
}

func TestSample2(t *testing.T) {
	k := 1
	P := []int{1, 1}
	runSample(t, k, P)
}

func TestSample3(t *testing.T) {
	k := 2
	P := []int{1, 2, 2, 1}
	runSample(t, k, P)
}
