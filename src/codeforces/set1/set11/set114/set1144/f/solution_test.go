package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect bool) {
	res, n, edges := process(bufio.NewReader(strings.NewReader(s)))

	if len(res) > 0 != expect {
		t.Fatalf("Sample expect %t, but got %s", expect, res)
	}
	if !expect {
		return
	}
	g := NewGraph(n, len(edges))

	for i, cur := range edges {
		u, v := cur[0]-1, cur[1]-1
		if res[i] == '1' {
			g.AddEdge(u, v)
		} else {
			g.AddEdge(v, u)
		}
	}

	var dfs func(u int, cnt int)
	dfs = func(u int, cnt int) {
		if cnt == 2 {
			t.Fatalf("Sample result %s, not correct", res)
		}
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			dfs(g.to[i], cnt+1)
		}
	}

	for i := 0; i < n; i++ {
		dfs(i, 0)
	}

}

func TestSample1(t *testing.T) {
	s := `6 5
1 5
2 1
1 4
3 1
6 1
`
	expect := true
	runSample(t, s, expect)
}
