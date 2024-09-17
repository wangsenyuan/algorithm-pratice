package main

import "testing"

func runSample(t *testing.T, n int, edges [][]int, questions [][]int, expect bool) {
	res := solve(n, edges, questions)

	if len(res) > 0 != expect {
		t.Fatalf("Sample expect %t, but got %v", expect, res)
	}

	if !expect {
		return
	}

	g := NewGraph(n, 2*n)

	for i, e := range edges {
		u, v := e[0]-1, e[1]-1
		g.AddEdge(u, v, i)
		g.AddEdge(v, u, i)
	}

	dep := make([]int, n)
	fa := make([][]int, n)
	fv := make([]int, n)
	var dfs func(p int, u int)

	dfs = func(p int, u int) {
		fa[u] = make([]int, H)
		fa[u][0] = p
		for i := 1; i < H; i++ {
			fa[u][i] = fa[fa[u][i-1]][i-1]
		}

		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				dep[v] = dep[u] + 1
				fv[v] = res[g.val[i]]
				dfs(u, v)
			}
		}
	}

	dfs(0, 0)

	lca := func(u int, v int) int {
		if dep[u] < dep[v] {
			u, v = v, u
		}
		for i := H - 1; i >= 0; i-- {
			if dep[u]-(1<<i) >= dep[v] {
				u = fa[u][i]
			}
		}
		if u == v {
			return u
		}
		for i := H - 1; i >= 0; i-- {
			if fa[u][i] != fa[v][i] {
				u = fa[u][i]
				v = fa[v][i]
			}
		}
		return fa[u][0]
	}

	check := func(a, b int) int {
		p := lca(a, b)

		res := inf

		for a != p {
			res = min(res, fv[a])
			a = fa[a][0]
		}

		for b != p {
			res = min(res, fv[b])
			b = fa[b][0]
		}

		return res
	}

	for _, cur := range questions {
		a, b, g := cur[0]-1, cur[1]-1, cur[2]
		if g != check(a, b) {
			t.Fatalf("Sample result %v, not correct", res)
		}
	}

}

func TestSample1(t *testing.T) {
	n := 4
	edges := [][]int{
		{1, 2},
		{3, 2},
		{3, 4},
	}
	questions := [][]int{
		{1, 2, 5},
		{1, 3, 3},
	}
	expect := true
	runSample(t, n, edges, questions, expect)
}

func TestSample2(t *testing.T) {
	n := 6
	edges := [][]int{
		{1, 2},
		{1, 6},
		{3, 1},
		{1, 5},
		{4, 1},
	}
	questions := [][]int{
		{6, 1, 3},
		{3, 4, 1},
		{6, 5, 2},
		{1, 2, 5},
	}
	expect := true
	runSample(t, n, edges, questions, expect)
}

func TestSample3(t *testing.T) {
	n := 6
	edges := [][]int{
		{1, 2},
		{1, 6},
		{3, 1},
		{1, 5},
		{4, 1},
	}
	questions := [][]int{
		{6, 1, 1},
		{3, 4, 3},
		{6, 5, 3},
		{1, 2, 4},
	}
	expect := false
	runSample(t, n, edges, questions, expect)
}

func TestSample4(t *testing.T) {
	n := 6
	edges := [][]int{
		{1, 2}, {2, 3}, {3, 4}, {4, 5}, {5, 6},
	}
	questions := [][]int{
		{1, 2, 6},
		{1, 3, 5},
		{1, 4, 4},
		{1, 5, 3},
		{1, 6, 2},
	}
	expect := true
	runSample(t, n, edges, questions, expect)
}

func TestSample5(t *testing.T) {
	n := 4
	edges := [][]int{
		{1, 2}, {2, 3}, {1, 4},
	}
	questions := [][]int{
		{3, 4, 2},
	}
	expect := true
	runSample(t, n, edges, questions, expect)
}
