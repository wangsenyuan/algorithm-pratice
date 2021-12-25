package main

import "testing"

func runSample(t *testing.T, n int, rgs [][]int, expect int) {
	res := solve(n, rgs)

	if res != expect {
		t.Errorf("Sample %d %v, expect %d, but got %d", n, rgs, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	rgs := [][]int{
		{1, 4},
		{2, 6},
		{5, 7},
	}
	expect := 1
	runSample(t, n, rgs, expect)
}

func TestSample2(t *testing.T) {
	n := 2
	rgs := [][]int{
		{1, 4},
		{4, 6},
	}
	expect := -1
	runSample(t, n, rgs, expect)
}

func TestSample3(t *testing.T) {
	n := 2
	rgs := [][]int{
		{3, 7},
		{8, 9},
	}
	expect := 0
	runSample(t, n, rgs, expect)
}
