package main

import "testing"

func runSample(t *testing.T, n int, P []int, rgs [][]int, expect int) {
	res := solve(n, P, rgs)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 2
	P := []int{1}
	rgs := [][]int{
		{1, 5},
		{2, 9},
	}
	expect := 1
	runSample(t, n, P, rgs, expect)
}

func TestSample2(t *testing.T) {
	n := 3
	P := []int{1, 1}
	rgs := [][]int{
		{4, 5},
		{2, 4},
		{6, 10},
	}
	expect := 2
	runSample(t, n, P, rgs, expect)
}

func TestSample3(t *testing.T) {
	n := 4
	P := []int{1, 2, 1}
	rgs := [][]int{
		{6, 9},
		{5, 6},
		{4, 5},
		{2, 4},
	}
	expect := 2
	runSample(t, n, P, rgs, expect)
}

func TestSample4(t *testing.T) {
	n := 5
	P := []int{1, 2, 3, 4}
	rgs := [][]int{
		{5, 5},
		{4, 4},
		{3, 3},
		{2, 2},
		{1, 1},
	}
	expect := 5
	runSample(t, n, P, rgs, expect)
}
