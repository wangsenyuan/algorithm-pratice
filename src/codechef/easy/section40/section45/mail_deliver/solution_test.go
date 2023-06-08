package main

import "testing"

func runSample(t *testing.T, n int, E [][]int, X []int, D []int, expect bool) {
	res := solve(n, E, X, D)

	if res != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	E := [][]int{
		{1, 2},
		{1, 3},
		{2, 3},
	}
	X := []int{1, 3}
	D := []int{1, 2}
	expect := true
	runSample(t, n, E, X, D, expect)
}

func TestSample2(t *testing.T) {
	n := 3
	E := [][]int{
		{1, 2},
		{1, 3},
		{2, 3},
	}
	X := []int{1}
	D := []int{1}
	expect := false
	runSample(t, n, E, X, D, expect)
}
