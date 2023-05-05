package main

import "testing"

func runSample(t *testing.T, n int, E [][]int, s int, x int, expect int) {
	res := solve(n, E, s, x)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 4
	E := [][]int{
		{1, 2},
		{3, 4},
		{2, 3},
		{2, 4},
	}
	s, x := 1, 4
	var expect = 2
	runSample(t, n, E, s, x, expect)
}

func TestSample2(t *testing.T) {
	n := 6
	E := [][]int{
		{1, 4},
		{1, 6},
		{1, 5},
		{1, 2},
		{5, 6},
		{4, 6},
		{6, 3},
		{2, 6},
	}
	s, x := 6, 1
	var expect = 4
	runSample(t, n, E, s, x, expect)
}

func TestSample3(t *testing.T) {
	n := 5
	E := [][]int{
		{3, 5},
		{5, 4},
		{3, 1},
		{4, 2},
		{2, 1},
		{1, 4},
	}
	s, x := 1, 3
	var expect = 1
	runSample(t, n, E, s, x, expect)
}

func TestSample4(t *testing.T) {
	n := 8
	E := [][]int{
		{2, 1},
		{3, 1},
		{4, 2},
		{5, 2},
		{6, 5},
		{7, 3},
		{8, 4},
		{6, 4},
		{8, 7},
		{1, 4},
		{4, 7},
		{1, 6},
		{6, 7},
		{3, 8},
		{8, 5},
		{4, 5},
		{4, 3},
		{8, 2},
	}
	s, x := 5, 1
	var expect = 11
	runSample(t, n, E, s, x, expect)
}

func TestSample5(t *testing.T) {
	n := 4
	E := [][]int{
		{1, 2},
		{2, 3},
		{1, 3},
		{3, 4},
	}
	s, x := 1, 4
	var expect = 2
	runSample(t, n, E, s, x, expect)
}
