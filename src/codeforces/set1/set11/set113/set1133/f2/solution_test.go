package main

import "testing"

func runSample(t *testing.T, n int, edges [][]int, D int, expect bool) {
	res := solve(n, edges, D)

	if len(res) == n-1 != expect {
		t.Fatalf("Sample expect %t, but got %v", expect, res)
	}

	if !expect {
		return
	}

	var deg int

	set := NewDSU(n + 1)

	for _, e := range res {
		u, v := e[0], e[1]
		if u == 1 || v == 1 {
			deg++
		}
		if !set.Union(u, v) {
			t.Fatalf("Sample result %v, not correct, detected an loop at %v", res, e)
		}
	}

	if deg != D {
		t.Fatalf("Sample result %v, not meeting the condition deg of 1 = D", res)
	}

	root := set.Find(1)
	if set.cnt[root] != n {
		t.Fatalf("Sample result %v, not correct, not all verts in the same component.", res)
	}
}

func TestSample1(t *testing.T) {
	n, D := 4, 1
	edges := [][]int{
		{1, 2},
		{1, 3},
		{1, 4},
		{2, 3},
		{3, 4},
	}
	expect := true
	runSample(t, n, edges, D, expect)
}

func TestSample2(t *testing.T) {
	n, D := 4, 3
	edges := [][]int{
		{1, 2},
		{1, 3},
		{1, 4},
		{2, 3},
		{3, 4},
	}
	expect := true
	runSample(t, n, edges, D, expect)
}

func TestSample3(t *testing.T) {
	n, D := 4, 3
	edges := [][]int{
		{1, 2},
		{1, 4},
		{2, 3},
		{3, 4},
	}
	expect := false
	runSample(t, n, edges, D, expect)
}

func TestSample4(t *testing.T) {
	n, D := 4, 2
	edges := [][]int{
		{4, 1},
		{2, 4},
		{1, 2},
		{1, 3},
		{4, 3},
	}
	expect := true
	runSample(t, n, edges, D, expect)
}
