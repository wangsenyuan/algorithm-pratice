package main

import (
	"bufio"
	"math/bits"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect bool) {
	res, n, edges := process(bufio.NewReader(strings.NewReader(s)))

	if len(res) > 0 != expect {
		t.Fatalf("Sample expect %t, but got %v", expect, res)
	}

	if !expect {
		return
	}

	g := NewGraph(n+1, 2*n)

	for _, e := range edges {
		u, v, w := e[0], e[1], e[2]
		g.AddEdge(u, v, w)
		g.AddEdge(v, u, w)
	}
	dep := make([]int, n+1)
	h := bits.Len(uint(n))
	fa := make([][]int, n+1)
	var dfs func(p int, u int)

	want := make([]int, n+1)

	dfs = func(p int, u int) {
		fa[u] = make([]int, h)
		fa[u][0] = p
		for i := 1; i < h; i++ {
			fa[u][i] = fa[fa[u][i-1]][i-1]
		}
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				dep[v] = dep[u] + 1
				want[v] = g.val[i]
				dfs(u, v)
			}
		}
	}
	fa[0] = make([]int, h)
	dfs(0, 1)

	lca := func(u int, v int) int {
		if dep[u] < dep[v] {
			u, v = v, u
		}
		for i := h - 1; i >= 0; i-- {
			if dep[u]-1<<i >= dep[v] {
				u = fa[u][i]
			}
		}
		if u == v {
			return u
		}
		for i := h - 1; i >= 0; i-- {
			if fa[u][i] != fa[v][i] {
				u = fa[u][i]
				v = fa[v][i]
			}
		}
		return fa[u][0]
	}

	add := make([]int, n+1)

	for _, op := range res {
		u, v, x := op[0], op[1], op[2]
		p := lca(u, v)
		add[u] += x
		add[v] += x
		add[fa[p][0]] -= x
	}

	var dfs2 func(p int, u int)

	dfs2 = func(p int, u int) {
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				dfs2(u, v)
				add[u] += add[v]
			}
		}
		if u > 1 && add[u] != want[u] {
			t.Fatalf("Sample result %v, not correct", res)
		}
	}

	dfs2(0, 1)
}

func TestSample1(t *testing.T) {
	s := `5
1 2 2
2 3 4
3 4 10
3 5 18
`
	expect := false
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `6
1 2 6
1 3 8
1 4 12
2 5 2
2 6 4
`
	expect := true
	runSample(t, s, expect)
}
