package main

import "testing"

func runSample(t *testing.T, n int, E [][]int, expect int) {
	edges := make([][]int, len(E))
	for i := 0; i < len(E); i++ {
		edges[i] = make([]int, 2)
		copy(edges[i], E[i])
	}
	cost, res := solve(n, edges)

	if cost != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, cost)
	}

	deg := make([]int, n)
	for _, e := range E {
		u, v := e[0]-1, e[1]-1
		deg[u]++
		deg[v]++
	}
	deg2 := make([]int, n)
	dsu := NewDSU(n)

	for _, e := range res {
		u, v := e[0]-1, e[1]-1
		deg2[u]++
		deg2[v]++
		dsu.Union(u, v)
	}

	x := dsu.Find(0)

	for i := 1; i < n; i++ {
		if dsu.Find(i) != x {
			t.Fatalf("Sample result %v, not correct", res)
		}
	}

	var num int

	for i := 0; i < n; i++ {
		num += abs(deg2[i] - deg[i])
	}

	if num != expect {
		t.Fatalf("Sample result %v, not correct, geting cost %d", res, num)
	}

}

func TestSample1(t *testing.T) {
	n := 3
	E := [][]int{
		{1, 2},
		{1, 3},
		{2, 3},
	}
	expect := 0
	runSample(t, n, E, expect)
}

func TestSample2(t *testing.T) {
	n := 3
	E := [][]int{
		{1, 2},
	}
	expect := 2
	runSample(t, n, E, expect)
}

func TestSample3(t *testing.T) {
	n := 4
	E := [][]int{
		{1, 2},
		{1, 2},
		{2, 3},
	}
	expect := 2
	runSample(t, n, E, expect)
}

func TestSample4(t *testing.T) {
	n := 5
	E := [][]int{
		{1, 2},
		{1, 2},
		{2, 3},
		{4, 5},
	}
	expect := 0
	runSample(t, n, E, expect)
}

func TestSample5(t *testing.T) {
	n := 6
	E := [][]int{
		{1, 2},
		{1, 2},
		{2, 3},
		{4, 5},
		{5, 6},
		{4, 6},
	}
	expect := 0
	runSample(t, n, E, expect)
}

func TestSample6(t *testing.T) {
	n := 7
	E := [][]int{
		{1, 2},
		{1, 2},
		{2, 3},
		{4, 5},
		{5, 6},
		{4, 6},
	}
	expect := 2
	runSample(t, n, E, expect)
}

func TestSample7(t *testing.T) {
	n := 11
	E := [][]int{
		{1, 2},
		{3, 4},
		{6, 7},
		{6, 7},
		{7, 8},
		{8, 9},
		{9, 10},
		{10, 11},
		{10, 11},
		{10, 11},
	}
	expect := 2
	runSample(t, n, E, expect)
}

func TestSample8(t *testing.T) {
	n := 13
	E := [][]int{
		{1, 2}, {1, 2}, {2, 3}, {3, 4}, {4, 1}, {5, 6}, {6, 7}, {7, 5}, {8, 9}, {10, 11},
	}
	expect := 4
	runSample(t, n, E, expect)
}
