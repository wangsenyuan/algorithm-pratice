package main

import "testing"

func runSample(t *testing.T, n int, E [][]int, k int, expect int) {
	res := solve(n, E, k)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	E := [][]int{
		{1, 2},
		{1, 3},
	}
	k := 2
	expect := 25
	runSample(t, n, E, k, expect)
}

func TestSample2(t *testing.T) {
	n := 7
	E := [][]int{
		{1, 2},
		{2, 3},
		{2, 4},
		{1, 5},
		{4, 6},
		{4, 7},
	}
	k := 2
	expect := 849
	runSample(t, n, E, k, expect)
}
