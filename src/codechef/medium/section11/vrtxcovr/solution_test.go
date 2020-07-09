package main

import "testing"

func runSample(t *testing.T, n, m int, E [][]int, expect bool) {
	ok, res := solve(n, m, E)

	if ok != expect {
		t.Errorf("Sample %d %d %v, expect %t, but got %t", n, m, E, expect, ok)
		return
	}

	if !expect {
		return
	}

	for i := 0; i < n; i += 2 {
		j := i ^ 1
		if res[i] == res[j] {
			t.Fatalf("Sample result %s, fail with %d and %d in same set", res, i, j)
		}
	}

	for i, e := range E {
		u, v := e[0], e[1]
		u--
		v--
		if res[u] == '0' && res[v] == '0' {
			t.Fatalf("Sample result %s, fail with edge %d, not covered", res, i)
		}
	}
}

func TestSample1(t *testing.T) {
	n, m := 3, 2
	E := [][]int{
		{1, 2},
		{2, 3},
	}
	expect := true

	runSample(t, n, m, E, expect)
}

func TestSample2(t *testing.T) {
	n, m := 4, 5
	E := [][]int{
		{1, 3},
		{2, 4},
		{1, 4},
		{1, 2},
		{2, 3},
	}
	expect := false

	runSample(t, n, m, E, expect)
}
