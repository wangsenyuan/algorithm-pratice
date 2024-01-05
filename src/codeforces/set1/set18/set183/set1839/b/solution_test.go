package main

import "testing"

func runSample(t *testing.T, lambs [][]int, expect int) {
	res := solve(lambs)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	lambs := [][]int{
		{2, 2},
		{1, 6},
		{1, 10},
		{1, 13},
	}
	expect := 15
	runSample(t, lambs, expect)
}

func TestSample2(t *testing.T) {
	lambs := [][]int{
		{3, 4},
		{3, 1},
		{2, 5},
		{3, 2},
		{3, 3},
	}
	expect := 14
	runSample(t, lambs, expect)
}

func TestSample3(t *testing.T) {
	lambs := [][]int{
		{1, 2},
		{3, 4},
		{1, 4},
		{3, 4},
		{3, 5},
		{2, 3},
	}
	expect := 20
	runSample(t, lambs, expect)
}
