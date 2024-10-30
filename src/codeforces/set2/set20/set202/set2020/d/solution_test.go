package main

import "testing"

func runSample(t *testing.T, n int, operations [][]int, expect int) {
	res := solve(n, operations)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 10
	operations := [][]int{
		{1, 2, 4},
		{2, 2, 4},
	}
	expect := 2
	runSample(t, n, operations, expect)
}

func TestSample2(t *testing.T) {
	n := 100
	operations := [][]int{
		{19, 2, 4},
	}
	expect := 96
	runSample(t, n, operations, expect)
}

func TestSample3(t *testing.T) {
	n := 100
	operations := [][]int{
		{1, 2, 5},
		{7, 2, 6},
		{17, 2, 31},
	}
	expect := 61
	runSample(t, n, operations, expect)
}
