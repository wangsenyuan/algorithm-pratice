package main

import (
	"testing"
)

func runSample(t *testing.T, n int, E [][]int, expect int) {
	res := solve(n, E)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 4
	E := [][]int{
		{1, 2},
		{2, 3},
		{3, 4},
	}
	expect := 1
	runSample(t, n, E, expect)
}

func TestSample2(t *testing.T) {
	n := 7
	E := [][]int{
		{1, 2},
		{1, 3},
		{1, 4},
		{1, 5},
		{1, 6},
		{1, 7},
	}
	expect := 0
	runSample(t, n, E, expect)
}

func TestSample3(t *testing.T) {
	n := 6
	E := [][]int{
		{1, 2},
		{2, 3},
		{2, 4},
		{5, 4},
		{6, 4},
	}
	expect := 4
	runSample(t, n, E, expect)
}

func TestSample4(t *testing.T) {
	n := 9
	E := [][]int{
		{4, 7},
		{5, 4},
		{2, 7},
		{5, 6},
		{3, 7},
		{7, 1},
		{9, 2},
		{8, 3},
	}
	expect := 8
	runSample(t, n, E, expect)
}
