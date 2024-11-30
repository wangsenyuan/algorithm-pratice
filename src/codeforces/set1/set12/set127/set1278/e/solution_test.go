package main

import (
	"slices"
	"testing"
)

func runSample(t *testing.T, n int, edges [][]int) {
	res := solve(n, edges)

	g := make([]map[int]bool, n)
	for i := range n {
		g[i] = make(map[int]bool)
	}

	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		g[u][v] = true
		g[v][u] = true
	}

	cnt := make([]int, 2*n+1)

	for _, cur := range res {
		cnt[cur.first]++
		cnt[cur.second]++
	}

	x := slices.Min(cnt[1:])
	y := slices.Max(cnt[1:])

	if x != y || x != 1 {
		t.Fatalf("Sample result %v, not correct", res)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}

			u, v := i, j
			if res[i].first > res[j].first {
				u, v = v, u
			}

			if res[u].first < res[v].second && res[v].first < res[u].second && res[u].second < res[v].second {
				if !g[i][j] {
					t.Fatalf("Sample result %v, not correct", res)
				}
			} else {
				if g[i][j] {
					t.Fatalf("Sample result %v, not correct", res)
				}
			}
		}
	}
}

func TestSample1(t *testing.T) {
	n := 6
	edges := [][]int{
		{1, 2},
		{1, 3},
		{3, 4},
		{3, 5},
		{2, 6},
	}
	runSample(t, n, edges)
}

func TestSample2(t *testing.T) {
	n := 1

	runSample(t, n, nil)
}
