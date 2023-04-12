package main

import "testing"

func runSample(t *testing.T, table [][]int, expect int64) {
	res := solve(table)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	table := [][]int{
		{1, 2, 3},
		{3, 2, 1},
	}
	var expect int64 = 7
	runSample(t, table, expect)
}

func TestSample2(t *testing.T) {
	table := [][]int{
		{1, 1, 2, 2},
		{2, 1, 1, 2},
		{2, 2, 1, 1},
	}
	var expect int64 = 76
	runSample(t, table, expect)
}

func TestSample3(t *testing.T) {
	table := [][]int{
		{1, 1, 2, 3},
		{2, 1, 1, 2},
		{3, 1, 2, 1},
		{1, 1, 2, 1},
	}
	var expect int64 = 129
	runSample(t, table, expect)
}
