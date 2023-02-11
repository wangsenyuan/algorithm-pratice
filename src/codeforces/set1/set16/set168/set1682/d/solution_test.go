package main

import "testing"

func runSample(t *testing.T, n int, S string, expect bool) {
	res := solve(n, S)

	if len(res) == n-1 != expect {
		t.Fatalf("Sample expect %t, but got %v", expect, res)
	}

	if !expect {
		return
	}

	g := make([]map[int]bool, n)
	for i := 0; i < n; i++ {
		g[i] = make(map[int]bool)
	}

	for _, x := range res {
		u, v := x[0], x[1]
		u--
		v--
		g[u][v] = true
		g[v][u] = true
	}

	for i := 0; i < n; i++ {
		x := int(S[i] - '0')
		if len(g[i])&1 != x {
			t.Fatalf("Sample result %v, not correct", res)
		}
	}

	vis := make([]bool, n)

	var dfs func(u int)
	dfs = func(u int) {
		vis[u] = true
		for v := range g[u] {
			if !vis[v] {
				dfs(v)
			}
		}
	}

	dfs(0)

	for i := 0; i < n; i++ {
		if !vis[i] {
			t.Fatalf("Sample result %v, not giveing a tree", res)
		}
	}
}

func TestSample1(t *testing.T) {
	n := 4
	S := "0110"
	expect := true

	runSample(t, n, S, expect)
}

func TestSample2(t *testing.T) {
	n := 2
	S := "10"
	expect := false

	runSample(t, n, S, expect)
}

func TestSample3(t *testing.T) {
	n := 6
	S := "110110"
	expect := true

	runSample(t, n, S, expect)
}
