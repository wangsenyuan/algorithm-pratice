package main

import "testing"

func runSample(t *testing.T, n int, E [][]int, P []int, expect int) {
	res := solve(n, E, P)
	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 4
	E := [][]int{
		{1, 2},
		{2, 3},
		{3, 4},
	}
	P := []int{2, 2}
	expect := 17
	runSample(t, n, E, P, expect)
}

func TestSample2(t *testing.T) {
	n := 4
	E := [][]int{
		{3, 4},
		{1, 3},
		{3, 2},
	}
	P := []int{3, 2}
	expect := 18
	runSample(t, n, E, P, expect)
}

func TestSample3(t *testing.T) {
	n := 7
	E := [][]int{
		{6, 1},
		{2, 3},
		{4, 6},
		{7, 3},
		{5, 1},
		{3, 6},
	}
	P := []int{7, 5, 13, 3}
	expect := 286
	runSample(t, n, E, P, expect)
}

func TestSample4(t *testing.T) {
	n := 16
	E := [][]int{
		{5, 10},
		{16, 1},
		{14, 1},
		{7, 5},
		{13, 2},
		{16, 11},
		{1, 7},
		{1, 4},
		{3, 14},
		{8, 16},
		{1, 6},
		{4, 9},
		{4, 12},
		{5, 13},
		{1, 15},
	}
	P := []int{45893, 9901, 51229, 15511, 46559, 28433, 4231, 30241,
		29837, 34421, 12953, 6577, 12143, 52711, 40493, 89, 34819, 28571}
	expect := 667739716
	runSample(t, n, E, P, expect)
}
