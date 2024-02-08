package main

import "testing"

func runSample(t *testing.T, n int, roads [][]int, a int, b int, expect bool) {
	res := solve(n, roads, a, b)

	if res != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, a, b := 3, 2, 1
	roads := [][]int{
		{2, 1},
		{3, 2},
		{1, 3},
	}
	expect := true
	runSample(t, n, roads, a, b, expect)
}

func TestSample2(t *testing.T) {
	n, a, b := 4, 1, 4
	roads := [][]int{
		{1, 4},
		{1, 2},
		{1, 3},
		{2, 3},
	}
	expect := false
	runSample(t, n, roads, a, b, expect)
}

func TestSample3(t *testing.T) {
	n, a, b := 4, 1, 2
	roads := [][]int{
		{1, 2},
		{2, 3},
		{2, 4},
		{3, 4},
	}
	expect := true
	runSample(t, n, roads, a, b, expect)
}

func TestSample4(t *testing.T) {
	n, a, b := 10, 6, 1
	roads := [][]int{
		{1, 2},
		{4, 3},
		{5, 8},
		{7, 8},
		{10, 4},
		{1, 9},
		{2, 4},
		{8, 1},
		{6, 2},
		{3, 1},
	}
	expect := true
	runSample(t, n, roads, a, b, expect)
}

func TestSample5(t *testing.T) {
	n, a, b := 7, 1, 1
	roads := [][]int{
		{4, 1},
		{2, 1},
		{5, 3},
		{4, 6},
		{4, 2},
		{7, 5},
		{3, 4},
	}
	expect := false
	runSample(t, n, roads, a, b, expect)
}
