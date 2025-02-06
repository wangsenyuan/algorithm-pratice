package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect bool) {
	reader := bufio.NewReader(strings.NewReader(s))

	n, edges, res := process(reader)

	if len(res) > 0 != expect {
		t.Fatalf("Sample expect %t, but got %v", expect, res)
	}

	if !expect {
		return
	}
	if len(res) != len(edges) {
		t.Fatalf("Sample result %v, not correct", res)
	}
	// 必须是一个强连通分量
	g := NewGraph(n, len(res))

	for _, cur := range res {
		u, v := cur[0]-1, cur[1]-1
		g.AddEdge(u, v, -1)
	}

	dsc := make([]int, n)
	low := make([]int, n)
	for i := 0; i < n; i++ {
		dsc[i] = -1
	}
	stack := make([]int, n)
	vis := make([]bool, n)
	var top int
	var timer int

	var dfs func(u int)
	dfs = func(u int) {
		dsc[u] = timer
		low[u] = timer
		timer++
		vis[u] = true
		stack[top] = u
		top++
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if vis[v] {
				low[u] = min(low[u], dsc[v])
			} else {
				dfs(v)
				low[u] = min(low[u], low[v])
			}
		}

		if low[u] == dsc[u] {
			for top > 0 {
				v := stack[top-1]
				top--
				vis[v] = false
				if u == v {
					break
				}
			}
		}
	}

	dfs(0)

	for i := 0; i < n; i++ {
		if dsc[i] < 0 {
			t.Fatalf("Sample result %v, not correct", res)
		}
	}

}

func TestSample1(t *testing.T) {
	s := `6 8
1 2
2 3
1 3
4 5
4 6
5 6
2 4
3 5
`
	expect := true
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `6 7
1 2
2 3
1 3
4 5
4 6
5 6
2 4
`
	expect := false
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `2 2
1 2
2 1
`
	expect := true
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `8 17
1 8
8 2
1 3
7 6
8 3
7 3
8 6
1 4
5 2
3 2
5 6
4 5
8 4
7 8
6 3
2 6
4 6
`
	expect := true
	runSample(t, s, expect)
}
