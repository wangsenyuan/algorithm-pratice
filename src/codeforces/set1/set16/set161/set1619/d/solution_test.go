package main

import "testing"

func runSample(t *testing.T, m, n int, P [][]int, expect int) {
	res := solve(m, n, P)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m := 2, 2
	P := [][]int{
		{1, 2},
		{3, 4},
	}
	expect := 3
	runSample(t, n, m, P, expect)
}

func TestSample2(t *testing.T) {
	m, n := 4, 3
	P := [][]int{
		{1, 3, 1},
		{3, 1, 1},
		{1, 2, 2},
		{1, 1, 3},
	}
	expect := 2
	runSample(t, m, n, P, expect)
}

func TestSample3(t *testing.T) {
	m, n := 2, 3
	P := [][]int{
		{5, 3, 4},
		{2, 5, 1},
	}
	expect := 4
	runSample(t, m, n, P, expect)
}

func TestSample4(t *testing.T) {
	m, n := 4, 2
	P := [][]int{
		{7, 9},
		{8, 1},
		{9, 6},
		{10, 8},
	}
	expect := 8
	runSample(t, m, n, P, expect)
}

func TestSample5(t *testing.T) {
	m, n := 2, 4
	P := [][]int{
		{6, 5, 2, 1},
		{7, 9, 7, 2},
	}
	expect := 2
	runSample(t, m, n, P, expect)
}
