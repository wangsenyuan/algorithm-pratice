package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, edges [][]int, cnt []int, expect bool) {
	label := solve(n, edges, cnt)

	if len(label) > 0 != expect {
		t.Fatalf("Sample expect %t, but got %v", expect, label)
	}

	if !expect {
		return
	}
	g := NewGraph(n, len(edges)*2)

	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	val := make([]int, 3)
	vis := make([]bool, n)

	var dfs func(u int)
	dfs = func(u int) {
		val[label[u]-1]++

		vis[u] = true

		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if abs(label[u]-label[v]) != 1 {
				t.Fatalf("Sample result %v not correct", label)
			}
			if !vis[v] {
				dfs(v)
			}
		}
	}

	for i := 0; i < n; i++ {
		if !vis[i] {
			dfs(i)
		}
	}

	if !reflect.DeepEqual(val, cnt) {
		t.Fatalf("Sample result %v, gives wrong number of cnt %v, expect %v", label, val, cnt)
	}
}

func abs(num int) int {
	return max(num, -num)
}

func TestSample1(t *testing.T) {
	n := 6
	edges := [][]int{
		{3, 1},
		{5, 4},
		{2, 5},
	}
	cnt := []int{2, 2, 2}
	expect := true
	runSample(t, n, edges, cnt, expect)
}

func TestSample2(t *testing.T) {
	n := 5
	edges := [][]int{
		{1, 2},
		{1, 3},
		{1, 5},
		{2, 3},
		{2, 4},
		{2, 5},
		{3, 4},
		{3, 5},
		{4, 5},
	}
	cnt := []int{0, 2, 3}
	expect := false
	runSample(t, n, edges, cnt, expect)
}
