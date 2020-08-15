package main

import "testing"

func runSample(t *testing.T, n int, m int, q int, E [][]int, expect int) {
	res := solve(n, m, q, E)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m, q := 7, 8, 25
	E := [][]int{
		{1, 2, 1},
		{2, 3, 10},
		{3, 4, 2},
		{1, 5, 2},
		{5, 6, 7},
		{6, 4, 15},
		{5, 3, 1},
		{1, 7, 3},
	}
	expect := 4361

	runSample(t, n, m, q, E, expect)
}

func TestSample2(t *testing.T) {
	n, m, q := 2, 1, 5
	E := [][]int{
		{1, 2, 4},
	}
	expect := 60

	runSample(t, n, m, q, E, expect)
}

func TestSample3(t *testing.T) {
	n, m, q := 15, 15, 23
	E := [][]int{
		{13, 10, 12},
		{11, 14, 12},
		{2, 15, 5},
		{4, 10, 8},
		{10, 2, 4},
		{10, 7, 5},
		{3, 10, 1},
		{5, 6, 11},
		{1, 13, 8},
		{9, 15, 4},
		{4, 2, 9},
		{11, 15, 1},
		{11, 12, 14},
		{10, 8, 12},
		{3, 6, 11},
	}
	expect := 3250

	runSample(t, n, m, q, E, expect)
}
func TestSample4(t *testing.T) {
	n, m, q := 5, 10, 10000000
	E := [][]int{
		{2, 4, 798},
		{1, 5, 824},
		{5, 2, 558},
		{4, 1, 288},
		{3, 4, 1890},
		{3, 1, 134},
		{2, 3, 1485},
		{4, 5, 284},
		{3, 5, 1025},
		{1, 2, 649},
	}
	expect := 768500592

	runSample(t, n, m, q, E, expect)
}
