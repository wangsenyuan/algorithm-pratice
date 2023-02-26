package main

import "testing"

func runSample(t *testing.T, n int, E [][]int) {
	vert, edge := solve(n, E)
	if len(vert) != n-1 {
		t.Fatalf("Sample should got %d nodes, but got %v", n, vert)
	}
	cnt := make([]int, n+1)
	for _, v := range vert {
		cnt[v[0]]++
		cnt[v[1]]++
	}

	for i := 1; i <= n; i++ {
		if cnt[i] == 0 {
			t.Fatalf("Sample result %v, not correct, as it don'est cover %d", vert, i)
		}
	}

	if len(edge) != n-2 {
		t.Fatalf("Sample result %v, not get a tree", edge)
	}
}

func TestSample1(t *testing.T) {
	n := 2
	E := [][]int{{1, 2}}
	runSample(t, n, E)
}

func TestSample2(t *testing.T) {
	n := 3
	E := [][]int{{1, 2}, {2, 3}}
	runSample(t, n, E)
}

func TestSample4(t *testing.T) {
	n := 4
	E := [][]int{
		{2, 1},
		{3, 1},
		{4, 1},
	}
	runSample(t, n, E)
}
