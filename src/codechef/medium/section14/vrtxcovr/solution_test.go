package main

import "testing"

func runSample(t *testing.T, n int, E [][]int, can bool) {
	ok, res := solve(n, E)

	if ok != can {
		t.Fatalf("Sample should get %t, but got %t", can, ok)
	}
	if !can {
		return
	}
	for _, e := range E {
		u, v := e[0], e[1]
		u--
		v--
		if res[u] == 0 && res[v] == 0 {
			t.Fatalf("Sample result %v, edge %v, neighter endpoint in set A", res, e)
		}
	}

	for i := 0; i < n; i++ {
		if i^1 < n && res[i] == res[i^1] {
			t.Fatalf("Sample result %v, %d & %d should not both in same set", res, i, i^1)
		}
	}
}

func TestSample1(t *testing.T) {
	n := 3
	E := [][]int{
		{1, 2},
		{2, 3},
	}
	can := true
	runSample(t, n, E, can)
}

func TestSample2(t *testing.T) {
	n := 4
	E := [][]int{
		{1, 3},
		{2, 4},
		{1, 4},
		{1, 2},
		{2, 3},
	}
	can := false
	runSample(t, n, E, can)
}
