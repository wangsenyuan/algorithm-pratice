package main

import "testing"

func runSample(t *testing.T, gas []int, edges [][]int, expect int) {
	res := solve(gas, edges)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	gas := []int{1, 3, 3}
	edges := [][]int{
		{1, 2, 2},
		{1, 3, 2},
	}
	expect := 3
	runSample(t, gas, edges, expect)
}

func TestSample2(t *testing.T) {
	gas := []int{6, 3, 2, 5, 0}
	edges := [][]int{
		{1, 2, 10},
		{2, 3, 3},
		{2, 4, 1},
		{1, 5, 1},
	}
	expect := 7
	runSample(t, gas, edges, expect)
}
