package main

import "testing"

func runSample(t *testing.T, n int, k int, lamps [][]int, expect int) {
	res := solve(n, k, lamps)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, k := 7, 3
	lamps := [][]int{
		{1, 7},
		{3, 8},
		{4, 5},
		{6, 7},
		{1, 3},
		{5, 10},
		{8, 9},
	}
	expect := 9
	runSample(t, n, k, lamps, expect)
}

func TestSample2(t *testing.T) {
	n, k := 3, 1
	lamps := [][]int{
		{1, 1},
		{2, 2},
		{3, 3},
	}
	expect := 3
	runSample(t, n, k, lamps, expect)
}

func TestSample3(t *testing.T) {
	n, k := 3, 3
	lamps := [][]int{
		{1, 3},
		{2, 3},
		{3, 3},
	}
	expect := 1
	runSample(t, n, k, lamps, expect)
}

func TestSample4(t *testing.T) {
	n, k := 5, 2
	lamps := [][]int{
		{1, 3},
		{2, 4},
		{3, 5},
		{4, 6},
		{5, 7},
	}
	expect := 7
	runSample(t, n, k, lamps, expect)
}

func TestSample5(t *testing.T) {
	n, k := 3, 2
	lamps := [][]int{
		{1, 1},
		{2, 2},
		{3, 3},
	}
	expect := 0
	runSample(t, n, k, lamps, expect)
}
