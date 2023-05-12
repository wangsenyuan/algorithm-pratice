package main

import "testing"

func runSample(t *testing.T, n int, F [][]int, expect int) {
	res := solve(n, F)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	F := [][]int{
		{3, 4},
		{5, 6},
		{5, 6},
	}
	expect := 1
	runSample(t, n, F, expect)
}

func TestSample2(t *testing.T) {
	n := 7
	F := [][]int{
		{3, 4},
		{5, 6},
		{5, 6},
		{7, 8},
		{7, 8},
		{13, 14},
		{14, 13},
	}
	expect := 3
	runSample(t, n, F, expect)
}
