package main

import "testing"

func runSample(t *testing.T, n int, k int, E [][]int, expect bool) {
	res := solve(n, k, E)

	if res != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	k := 2
	E := [][]int{
		{1, 4},
		{2, 6},
		{2, 5},
		{3, 5},
		{3, 6},
	}
	expect := false

	runSample(t, n, k, E, expect)
}

func TestSample2(t *testing.T) {
	n := 3
	k := 1
	E := [][]int{
		{1, 4},
		{2, 6},
		{2, 5},
		{3, 5},
		{3, 6},
	}
	expect := true

	runSample(t, n, k, E, expect)
}
