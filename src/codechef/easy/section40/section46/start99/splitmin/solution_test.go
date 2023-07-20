package main

import "testing"

func runSample(t *testing.T, pairs [][]int, expect int) {
	res := solve(pairs)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	pairs := [][]int{
		{1, 6},
		{4, 9},
	}
	expect := 2
	runSample(t, pairs, expect)
}

func TestSample2(t *testing.T) {
	pairs := [][]int{
		{1, 9},
		{4, 6},
	}
	expect := 3
	runSample(t, pairs, expect)
}

func TestSample3(t *testing.T) {
	pairs := [][]int{
		{10, 10},
		{23, 55},
		{21, 80},
		{105, 91},
	}
	expect := 11
	runSample(t, pairs, expect)
}
