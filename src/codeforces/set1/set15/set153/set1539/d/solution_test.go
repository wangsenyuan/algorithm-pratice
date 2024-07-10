package main

import "testing"

func runSample(t *testing.T, products [][]int, expect int) {
	res := solve(products, len(products))

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	products := [][]int{
		{3, 4},
		{1, 3},
		{1, 5},
	}
	expect := 8
	runSample(t, products, expect)
}

func TestSample2(t *testing.T) {
	products := [][]int{
		{2, 7},
		{2, 8},
		{1, 2},
		{2, 4},
		{1, 8},
	}
	expect := 12
	runSample(t, products, expect)
}
