package main

import "testing"

func runSample(t *testing.T, n int, buses [][]int, expect int) {
	res := solve(n, buses)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 2
	buses := [][]int{
		{0, 1},
		{1, 2},
	}
	expect := 1
	runSample(t, n, buses, expect)
}

func TestSample2(t *testing.T) {
	n := 5
	buses := [][]int{
		{0, 1},
		{0, 2},
		{0, 3},
		{0, 4},
		{0, 5},
	}
	expect := 16
	runSample(t, n, buses, expect)
}

func TestSample3(t *testing.T) {
	n := 3
	buses := [][]int{
		{1, 2},
		{2, 3},
		{1, 3},
	}
	expect := 0
	runSample(t, n, buses, expect)
}

func TestSample4(t *testing.T) {
	n := 3
	buses := [][]int{
		{0, 1},
		{1, 3},
		{2, 3},
		{4, 6},
		{5, 7},
		{4, 5},
		{5, 7},
	}
	expect := 0
	runSample(t, n, buses, expect)
}
