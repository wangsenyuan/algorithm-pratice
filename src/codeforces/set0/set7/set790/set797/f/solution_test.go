package main

import "testing"

func runSample(t *testing.T, x []int, holes [][]int, expect int) {
	res := solve(x, holes)
	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	x := []int{6, 2, 8, 9}
	holes := [][]int{
		{3, 6},
		{2, 1},
		{3, 6},
		{4, 7},
		{4, 7},
	}
	expect := 11
	runSample(t, x, holes, expect)
}

func TestSample2(t *testing.T) {
	x := []int{10, 20, 30, 40, 50, 45, 35}
	holes := [][]int{
		{-1000000000, 10},
		{1000000000, 1},
	}
	expect := 7000000130
	runSample(t, x, holes, expect)
}
