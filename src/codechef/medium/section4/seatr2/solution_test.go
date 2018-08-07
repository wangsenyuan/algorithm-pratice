package main

import "testing"

func runSample(t *testing.T, N, M int, edges [][]int, expect int) {
	res := solve(N, M, edges)
	if res != expect {
		t.Errorf("Sample %d %d %v, expect %d, but got %d", N, M, edges, expect, res)
	}
}

func TestSample1(t *testing.T) {
	N, M := 3, 2
	edges := [][]int{{1, 2}, {1, 3}}
	expect := 5
	runSample(t, N, M, edges, expect)
}

func TestSample2(t *testing.T) {
	N, M := 1, 100
	// edges := [][]int{{1, 2}, {1, 3}}
	expect := 100
	runSample(t, N, M, nil, expect)
}
