package main

import "testing"

func runSample(t *testing.T, n int, A []int, queries []int, expect []bool) {
	solver := NewSolver(n, A)

	for i := 0; i < len(queries); i++ {
		x := queries[i]

		res := solver.Add(x)

		if res != expect[i] {
			t.Fatalf("Sample fail at %d, expect %t, but got %t", i, expect[i], res)
		}
	}
}

func TestSample1(t *testing.T) {
	n := 6
	A := []int{1, 1, 1, 2, 1, 1}
	queries := []int{
		2, 1, -1, 2, -1, 2,
	}
	expect := []bool{false, true, false, false, false, true}

	runSample(t, n, A, queries, expect)
}

func TestSample2(t *testing.T) {
	n := 10
	A := []int{5, 1, 5, 1, 4, 4, 2, 2, 4, 4}
	queries := []int{
		-1,
		-4,
		+2,
		+2,
		-5,
		-5,
		-4,
		-4,
		+10,
		+10,
		+10,
		+10,
		-2,
		+1,
		-4,
	}
	expect := []bool{true, false, false, true, false, false, false, false, false, false, false, true, false, true, true}

	runSample(t, n, A, queries, expect)
}
