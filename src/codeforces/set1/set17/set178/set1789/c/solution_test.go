package main

import "testing"

func runSample(t *testing.T, a []int, ops [][]int, expect int) {
	res := solve(a, ops)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{1, 2, 3}
	ops := [][]int{
		{1, 4},
		{2, 5},
	}
	expect := 13
	runSample(t, a, ops, expect)
}

func TestSample2(t *testing.T) {
	a := []int{1}
	ops := [][]int{
		{1, 1},
	}
	expect := 1
	runSample(t, a, ops, expect)
}

func TestSample3(t *testing.T) {
	a := []int{4, 6, 9, 12, 16, 20, 2, 10, 19, 7}
	ops := [][]int{
		{1, 3},
		{5, 4},
		{2, 17},
		{2, 18},
		{6, 11},
		{7, 1},
		{8, 17},
		{5, 5},
		{5, 5},
		{2, 2},
	}
	expect := 705
	runSample(t, a, ops, expect)
}
