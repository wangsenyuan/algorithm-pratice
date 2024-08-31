package main

import "testing"

func runSample(t *testing.T, a [][]int, expect int) {
	res := solve(a)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	}
	expect := 9
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := [][]int{
		{2, 5, 4, 8, 3},
		{9, 10, 11, 5, 1},
		{12, 8, 4, 2, 5},
		{2, 2, 5, 4, 1},
		{6, 8, 2, 4, 2},
	}
	expect := 49
	runSample(t, a, expect)
}

func TestSample3(t *testing.T) {
	a := [][]int{
		{123456789876543, 987654321234567},
	}
	expect := 864197531358023
	runSample(t, a, expect)
}
