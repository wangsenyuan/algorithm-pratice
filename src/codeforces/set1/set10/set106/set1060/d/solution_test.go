package main

import "testing"

func runSample(t *testing.T, guests [][]int, expect int) {
	res := solve(guests)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	guests := [][]int{
		{1, 1},
		{1, 1},
		{1, 1},
	}
	expect := 6
	runSample(t, guests, expect)
}

func TestSample2(t *testing.T) {
	guests := [][]int{
		{1, 2},
		{2, 1},
		{3, 5},
		{5, 3},
	}
	expect := 15
	runSample(t, guests, expect)
}

func TestSample3(t *testing.T) {
	guests := [][]int{
		{5, 6},
	}
	expect := 7
	runSample(t, guests, expect)
}
