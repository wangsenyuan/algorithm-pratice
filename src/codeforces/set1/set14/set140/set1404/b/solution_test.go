package main

import "testing"

func runSample(t *testing.T, n int, a, b, da, db int, E [][]int, expect bool) {
	res := solve(n, a, b, da, db, E)

	if res != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, a, b, da, db := 4, 3, 2, 1, 2
	E := [][]int{
		{1, 2},
		{1, 3},
		{1, 4},
	}
	expect := true
	runSample(t, n, a, b, da, db, E, expect)
}

func TestSample2(t *testing.T) {
	n, a, b, da, db := 6, 6, 1, 2, 5
	E := [][]int{
		{1, 2},
		{6, 5},
		{2, 3},
		{3, 4},
		{4, 5},
	}
	expect := false
	runSample(t, n, a, b, da, db, E, expect)
}
