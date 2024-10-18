package main

import "testing"

func runSample(t *testing.T, n int, roads [][]int, expect bool) {
	ok, res := solve(n, roads)

	if ok != expect {
		t.Fatalf("Sample expect %t, but got %t, %v", expect, ok, res)
	}

	if !expect {
		return
	}

	ord := make([]bool, n+1)

	for _, u := range res {
		ord[u] = true
	}

	for _, road := range roads {
		u, v, w := road[0], road[1], road[2]
		if ord[u] {
			w ^= 1
		}
		if ord[v] {
			w ^= 1
		}
		if w != 1 {
			t.Fatalf("Sample result %v, not correct", res)
		}
	}
}

func TestSample1(t *testing.T) {
	n := 4
	edges := [][]int{
		{1, 2, 1},
		{2, 4, 0},
		{4, 3, 1},
		{3, 2, 0},
	}
	expect := true
	runSample(t, n, edges, expect)
}

func TestSample2(t *testing.T) {
	n := 3
	edges := [][]int{
		{1, 2, 0},
		{2, 3, 0},
		{3, 1, 0},
	}
	expect := false
	runSample(t, n, edges, expect)
}
