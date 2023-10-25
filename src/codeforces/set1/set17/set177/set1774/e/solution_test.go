package main

import "testing"

func runSample(t *testing.T, n int, E [][]int, a []int, b []int, d int, expect int) {
	res := solve(n, E, a, b, d)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, d := 4, 2
	E := [][]int{
		{1, 2},
		{1, 3},
		{2, 4},
	}
	a := []int{3}
	b := []int{4}
	expect := 6
	runSample(t, n, E, a, b, d, expect)
}

func TestSample2(t *testing.T) {
	n, d := 4, 2
	E := [][]int{
		{1, 2},
		{2, 3},
		{3, 4},
	}
	a := []int{1, 2, 3, 4}
	b := []int{1}
	expect := 8
	runSample(t, n, E, a, b, d, expect)
}
