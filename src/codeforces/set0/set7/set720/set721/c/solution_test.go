package main

import "testing"

func runSample(t *testing.T, n int, edges [][]int, maxDist int, expect int) {
	res := solve(n, edges, maxDist)

	if len(res) != expect {
		t.Fatalf("Sample expect %d length path, but got %v", expect, res)
	}

	if res[0] != 1 || res[len(res)-1] != n {
		t.Fatalf("path should from 1 to n but got %d -> %d", res[0], res[len(res)-1])
	}

	g := make([]map[int]int, n+1)
	for i := 0; i <= n; i++ {
		g[i] = make(map[int]int)
	}

	for _, e := range edges {
		u, v, w := e[0], e[1], e[2]
		g[u][v] = w
	}

	var dist int
	for i := 1; i < len(res); i++ {
		if d, ok := g[res[i-1]][res[i]]; !ok {
			t.Fatalf("Sample result has a non-exists edge between %d -> %d", res[i-1], res[i])
		} else {
			dist += d
		}
	}

	if dist > maxDist {
		t.Fatalf("Sample result %v, distance %d exceeds %d", res, dist, maxDist)
	}
}

func TestSample1(t *testing.T) {
	n, maxDist := 4, 13
	edges := [][]int{
		{1, 2, 5},
		{2, 3, 7},
		{2, 4, 8},
	}
	expect := 3
	runSample(t, n, edges, maxDist, expect)
}

func TestSample2(t *testing.T) {
	n, maxDist := 6, 7
	edges := [][]int{
		{1, 2, 2},
		{1, 3, 3},
		{3, 6, 3},
		{2, 4, 2},
		{4, 6, 2},
		{6, 5, 1},
	}
	expect := 4
	runSample(t, n, edges, maxDist, expect)
}

func TestSample3(t *testing.T) {
	n, maxDist := 5, 6
	edges := [][]int{
		{1, 3, 3},
		{3, 5, 3},
		{1, 2, 2},
		{2, 4, 3},
		{4, 5, 2},
	}
	expect := 3
	runSample(t, n, edges, maxDist, expect)
}
