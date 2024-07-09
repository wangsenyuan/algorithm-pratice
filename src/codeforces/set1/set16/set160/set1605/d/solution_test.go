package main

import "testing"

func runSample(t *testing.T, n int, edges [][]int, expect []int) {
	res := solve(n, edges)

	a := check(n, edges, expect)
	b := check(n, edges, res)

	if a != b {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func check(n int, edges [][]int, p []int) int {
	var cnt int
	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		a, b := p[u], p[v]
		if a^b > min(a, b) {
			cnt++
		}
	}

	return cnt
}

func TestSample1(t *testing.T) {
	n := 3
	edges := [][]int{
		{1, 2},
		{1, 3},
	}
	expect := []int{1, 2, 3}
	runSample(t, n, edges, expect)
}
