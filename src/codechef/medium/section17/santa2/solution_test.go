package main

import "testing"

func runSample(t *testing.T, n int, E [][]int, P []int, expect bool) {
	res := solve(n, E, P)

	if res != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	E := [][]int{
		{1, 2},
		{1, 3},
	}
	P := []int{1, 2, 3}
	expect := true
	runSample(t, n, E, P, expect)
}

func TestSample2(t *testing.T) {
	n := 3
	E := [][]int{
		{1, 2},
		{2, 3},
	}
	P := []int{1, 3, 2}
	expect := false
	runSample(t, n, E, P, expect)
}

func TestSample3(t *testing.T) {
	n := 3
	E := [][]int{
		{1, 2},
		{1, 3},
		{2, 3},
	}
	P := []int{1, 2, 3}
	expect := true
	runSample(t, n, E, P, expect)
}
