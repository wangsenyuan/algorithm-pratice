package main

import "testing"

func runSample(t *testing.T, n int, people [][]int, expect int) {
	res := solve(n, people)
	if res != expect {
		t.Errorf("Sample %d %v, expect %d, but got %d", n, people, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 2
	people := [][]int{
		{3, 5},
		{8, 100},
	}
	expect := 4
	runSample(t, n, people, expect)
}

func TestSample2(t *testing.T) {
	n := 3
	people := [][]int{
		{5, 100, 1},
		{2},
		{5, 100},
	}
	expect := 4
	runSample(t, n, people, expect)
}
