package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	f, res := process(bufio.NewReader(strings.NewReader(s)))

	if len(res) != expect {
		t.Fatalf("Sample expect %d, but got %v", expect, res)
	}
	n := len(f)
	g := NewGraph(n, n+len(res))
	for i := 0; i < n; i++ {
		g.AddEdge(i, f[i]-1)
	}

	for _, cur := range res {
		u, v := cur[0], cur[1]
		u--
		v--
		g.AddEdge(u, v)
	}

	vis := make([]bool, n)
	que := make([]int, n)
	bfs := func(x int) bool {
		clear(vis)
		var head, tail int
		que[head] = x
		head++
		vis[x] = true
		for tail < head {
			u := que[tail]
			tail++
			for i := g.nodes[u]; i > 0; i = g.next[i] {
				v := g.to[i]
				if !vis[v] {
					vis[v] = true
					que[head] = v
					head++
				}
			}
		}
		for i := 0; i < n; i++ {
			if !vis[i] {
				return false
			}
		}
		return true
	}
	for i := 0; i < n; i++ {
		if !bfs(i) {
			t.Fatalf("Sample result %v, not correct", res)
		}
	}
}

func TestSample1(t *testing.T) {
	s := `3
3 3 2
`
	expect := 1
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `7
2 3 1 3 4 4 1
`
	expect := 3
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `20
20 10 16 14 9 20 6 20 14 19 17 13 16 13 14 8 8 8 8 19
`
	expect := 10
	runSample(t, s, expect)
}
