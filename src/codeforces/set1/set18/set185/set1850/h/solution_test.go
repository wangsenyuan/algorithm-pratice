package main

import "testing"

func runSample(t *testing.T, n int, conditions [][]int, expect bool) {
	res := solve(n, conditions)

	if res != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 5
	conditions := [][]int{
		{1, 2, 2},
		{2, 3, 4},
		{4, 2, -6},
	}
	expect := true
	runSample(t, n, conditions, expect)
}

func TestSample2(t *testing.T) {
	n := 6
	conditions := [][]int{
		{1, 2, 2},
		{2, 3, 4},
		{4, 2, -6},
		{5, 4, 4},
		{3, 5, 100},
	}
	expect := false
	runSample(t, n, conditions, expect)
}

func TestSample3(t *testing.T) {
	n := 4
	conditions := [][]int{
		{3, 1, 3},
		{4, 2, 0},
		{1, 3, -3},
		{2, 3, -4},
	}
	expect := true
	runSample(t, n, conditions, expect)
}

func TestSample4(t *testing.T) {
	n := 5
	conditions := [][]int{
		{3, 4, -7},
		{5, 1, -3},
		{3, 5, 2},
		{3, 2, 1},
		{5, 3, -2},
	}
	expect := true
	runSample(t, n, conditions, expect)
}
