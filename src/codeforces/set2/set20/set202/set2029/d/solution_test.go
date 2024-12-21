package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	res, n, edges := process(bufio.NewReader(strings.NewReader(s)))

	if len(res) > 2*max(n, len(edges)) {
		t.Fatalf("Sample result %v, took too much", res)
	}

	g := make([]map[int]bool, n)

	for i := 0; i < n; i++ {
		g[i] = make(map[int]bool)
	}
	add := func(u int, v int) {
		g[u][v] = true
		g[v][u] = true
	}
	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		add(u, v)
	}

	remove := func(u int, v int) {
		delete(g[u], v)
		delete(g[v], u)
	}

	update := func(u int, v int) {
		if g[u][v] {
			remove(u, v)
		} else {
			add(u, v)
		}
	}

	for _, cur := range res {
		u, v, w := cur[0]-1, cur[1]-1, cur[2]-1
		update(u, v)
		update(u, w)
		update(v, w)
	}

	if len(g[0]) == 0 {
		for i := 0; i < n; i++ {
			if len(g[i]) > 0 {
				t.Fatalf("Sample result %v, not correct", res)
			}
		}
		return
	}

	vis := make([]bool, n)

	var dfs func(p int, u int) bool

	dfs = func(p int, u int) bool {
		if vis[u] {
			return false
		}
		vis[u] = true
		for v := range g[u] {
			if p == v {
				continue
			}
			if !dfs(u, v) {
				return false
			}
		}
		return true
	}

	if !dfs(-1, 0) {
		t.Fatalf("Sample result %v, not correct", res)
	}

	for i := 0; i < n; i++ {
		if !vis[i] {
			t.Fatalf("Sample result %v, not correct", res)
		}
	}
}

func TestSample1(t *testing.T) {
	s := `3 0`
	runSample(t, s)
}

func TestSample2(t *testing.T) {
	s := `3 1
1 2`
	runSample(t, s)
}

func TestSample3(t *testing.T) {
	s := `3 2
1 2
2 3`
	runSample(t, s)
}

func TestSample4(t *testing.T) {
	s := `3 3
1 2
2 3
3 1`
	runSample(t, s)
}

func TestSample5(t *testing.T) {
	s := `6 6
1 2
1 6
4 5
3 4
4 6
3 6`
	runSample(t, s)
}
