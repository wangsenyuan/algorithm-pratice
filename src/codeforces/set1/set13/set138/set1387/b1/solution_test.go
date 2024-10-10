package main

import "testing"

func runSample(t *testing.T, n int, edges [][]int, expect int) {
	res, steps := solve(n, edges)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}

	dist := make([][]int, n+1)

	for i := 0; i <= n; i++ {
		dist[i] = make([]int, n+1)
		for j := 0; j <= n; j++ {
			dist[i][j] = inf
		}
		dist[i][i] = 0
	}

	for _, e := range edges {
		u, v := e[0], e[1]
		dist[u][v] = 1
		dist[v][u] = 1
	}

	for k := 1; k <= n; k++ {
		for i := 1; i <= n; i++ {
			for j := 1; j <= n; j++ {
				dist[i][j] = min(dist[i][j], dist[i][k]+dist[k][j])
			}
		}
	}

	var sum int

	for u := 1; u <= n; u++ {
		v := steps[u-1]
		if u == v {
			t.Fatalf("Sample result %v, %d stay at the original house, now allowed", res, u)
		}
		sum += dist[u][v]
	}

	if sum != expect {
		t.Fatalf("Sample result %v, not correct", steps)
	}
}

func TestSample1(t *testing.T) {
	n := 4
	edges := [][]int{
		{1, 2},
		{2, 3},
		{3, 4},
	}
	expect := 4
	runSample(t, n, edges, expect)
}

func TestSample2(t *testing.T) {
	n := 7
	edges := [][]int{
		{4, 2},
		{5, 7},
		{3, 4},
		{6, 3},
		{1, 3},
		{4, 5},
	}
	expect := 8
	runSample(t, n, edges, expect)
}
