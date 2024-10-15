package main

import "testing"

func runSample(t *testing.T, n int, edges [][]int, expect int) {
	res := solve(n, edges)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 14
	edges := [][]int{
		{1, 2},
		{2, 7},
		{3, 4},
		{6, 3},
		{5, 7},
		{3, 8},
		{6, 8},
		{11, 12},
	}
	expect := 1
	runSample(t, n, edges, expect)
}

func TestSample2(t *testing.T) {
	n := 200000
	edges := [][]int{
		{7, 9},
		{9, 8},
		{4, 5},
	}
	expect := 0
	runSample(t, n, edges, expect)
}

func TestSample3(t *testing.T) {
	n := 500
	edges := [][]int{
		{100, 300},
		{200, 400},
		{420, 440},
		{430, 450},
		{435, 460},
	}
	expect := 335
	runSample(t, n, edges, expect)
}
