package main

import "testing"

func runSample(t *testing.T, n int, E [][]int, A []int) {
	res := solve(n, E)

	expect := process(n, E, A)
	ans := process(n, E, A)

	if expect != ans {
		t.Fatalf("Sample result %v, not giving correct min answer %d, but got %d", res, expect, ans)
	}
}

func process(n int, E [][]int, A []int) int64 {
	var sum int64
	for _, e := range E {
		u, v, w := e[0]-1, e[1]-1, e[2]
		b := A[u] - A[v]
		// b > 0
		sum += int64(b) * int64(w)
	}
	return sum
}

func TestSample1(t *testing.T) {
	n := 3
	E := [][]int{
		{2, 1, 4},
		{1, 3, 2},
	}
	A := []int{1, 2, 0}
	runSample(t, n, E, A)
}

func TestSample2(t *testing.T) {
	n := 5
	E := [][]int{
		{1, 2, 1},
		{2, 3, 1},
		{1, 3, 6},
		{4, 5, 8},
	}
	A := []int{43, 42, 41, 1337, 1336}
	runSample(t, n, E, A)
}

func TestSample3(t *testing.T) {
	n := 5
	E := [][]int{
		{1, 2, 1},
		{2, 3, 1},
		{3, 4, 1},
		{1, 5, 1},
		{5, 4, 10},
	}
	A := []int{4, 3, 2, 1, 2}
	runSample(t, n, E, A)
}
