package main

import "testing"

func runSample(t *testing.T, n int, edges [][]int, expect int) {
	res := solve(n, edges)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 2
	edges := [][]int{
		{1, 2},
	}
	expect := 4
	runSample(t, n, edges, expect)
}
func TestSample2(t *testing.T) {
	n := 4
	edges := [][]int{
		{1, 2},
		{1, 3},
		{1, 4},
		{2, 3},
		{2, 4},
		{3, 4},
	}
	expect := 0
	runSample(t, n, edges, expect)
}

func TestSample3(t *testing.T) {
	n := 100000
	expect := 296595689
	runSample(t, n, nil, expect)
}
