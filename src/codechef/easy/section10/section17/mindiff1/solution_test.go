package main

import "testing"

func runSample(t *testing.T, n int, E [][]int, expect int) {
	res, C := solve(n, E)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	} else {
		D := make([]int, n)
		for _, e := range E {
			u, v := e[0], e[1]
			u--
			v--
			if C[u] < C[v] {
				D[v]++
			} else {
				D[u]++
			}
		}

		for i := 0; i < n; i++ {
			if D[i] > res {
				t.Fatalf("Sample result %v, not correct", C)
			}
		}

	}
}

func TestSample1(t *testing.T) {
	n := 5
	E := [][]int{
		{1, 2},
		{1, 3},
		{1, 4},
		{2, 3},
		{2, 4},
		{2, 5},
		{3, 5},
	}
	expect := 2
	runSample(t, n, E, expect)
}

func TestSample2(t *testing.T) {
	n := 5
	E := [][]int{
		{1, 2},
		{2, 3},
		{3, 4},
		{4, 5},
	}
	expect := 1
	runSample(t, n, E, expect)
}

func TestSample3(t *testing.T) {
	n := 5
	E := [][]int{
		{1, 2},
		{2, 3},
		{1, 3},
	}
	expect := 2
	runSample(t, n, E, expect)
}
