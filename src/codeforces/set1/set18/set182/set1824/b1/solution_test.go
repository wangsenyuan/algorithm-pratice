package main

import "testing"

func runSample(t *testing.T, n int, routes [][]int, k int, expect int) {
	res := solve(n, routes, k)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, k := 4, 2
	routes := [][]int{
		{1, 2},
		{2, 3},
		{3, 4},
	}
	expect := 666666674
	runSample(t, n, routes, k, expect)
}
