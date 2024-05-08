package main

import "testing"

func runSample(t *testing.T, n int, h int, u []int, edges [][]int, expect int) {
	res := solve(n, h, u, edges)

	if len(res) != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, len(res))
	}

	// update time
	for _, i := range res {
		u[i-1]++
		if u[i-1] == h {
			u[i-1] = 0
		}
	}

	// check
	for _, e := range edges {
		a, b := e[0]-1, e[1]-1
		if u[a] == u[b] {
			t.Fatalf("Sample result %v, causing conflict", res)
		}
	}
}

func TestSample1(t *testing.T) {
	n, h := 3, 5
	u := []int{4, 4, 0}
	edges := [][]int{
		{1, 3},
		{3, 2},
		{3, 1},
	}
	expect := 1
	runSample(t, n, h, u, edges, expect)
}

func TestSample2(t *testing.T) {
	n, h := 4, 4
	u := []int{2, 1, 0, 3}
	edges := [][]int{
		{4, 3},
		{3, 2},
		{1, 2},
		{1, 4},
		{1, 3},
	}
	expect := 4
	runSample(t, n, h, u, edges, expect)
}

func TestSample3(t *testing.T) {
	n, h := 2, 2
	u := []int{1, 0}
	edges := [][]int{
		{1, 2},
	}
	expect := 2
	runSample(t, n, h, u, edges, expect)
}
