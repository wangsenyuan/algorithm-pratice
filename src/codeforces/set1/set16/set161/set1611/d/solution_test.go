package main

import "testing"

func runSample(t *testing.T, n int, B []int, P []int, expect bool) {
	res := solve(n, B, P)

	if len(res) > 0 != expect {
		t.Fatalf("Sample expect %t, but got %v", expect, res)
	}
	if !expect {
		return
	}

	g := NewGraph(n, n)
	for i := 0; i < n; i++ {
		p := B[i]
		p--
		if p != i {
			g.AddEdge(p, i)
		}
	}

	dist := make([]int, n)

	var dfs func(u int)

	dfs = func(u int) {
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			dist[v] = dist[u] + res[v]
			dfs(v)
		}
	}

	for i := 0; i < n; i++ {
		if B[i]-1 == i {
			dfs(i)
			break
		}
	}

	for i := 0; i+1 < n; i++ {
		if dist[P[i]-1] >= dist[P[i+1]-1] {
			t.Fatalf("Sample result %v, not correct", res)
		}
	}

}

func TestSample1(t *testing.T) {
	n := 5
	B := []int{3, 1, 3, 3, 1}
	P := []int{3, 1, 2, 5, 4}
	expect := true
	runSample(t, n, B, P, expect)
}

func TestSample2(t *testing.T) {
	n := 3
	B := []int{1, 1, 2}
	P := []int{3, 1, 2}
	expect := false
	runSample(t, n, B, P, expect)
}

func TestSample3(t *testing.T) {
	n := 7
	B := []int{1, 1, 2, 3, 4, 5, 6}
	P := []int{1, 2, 3, 4, 5, 6, 7}
	expect := true
	runSample(t, n, B, P, expect)
}

func TestSample4(t *testing.T) {
	n := 6
	B := []int{4, 4, 4, 4, 1, 1}
	P := []int{4, 2, 1, 5, 6, 3}
	expect := true
	runSample(t, n, B, P, expect)
}
