package main

import "testing"

func runSample(t *testing.T, n, m int, G [][]int, expect int) {
	res := solve(n, m, G)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m := 2, 2
	G := [][]int{
		{1, 1},
		{0, 1},
	}
	expect := 0
	runSample(t, n, m, G, expect)
}

func TestSample2(t *testing.T) {
	n, m := 2, 3
	G := [][]int{
		{1, 1, 0},
		{1, 0, 0},
	}
	expect := 3
	runSample(t, n, m, G, expect)
}

func TestSample3(t *testing.T) {
	n, m := 3, 7
	G := [][]int{
		{1, 0, 1, 1, 1, 1, 1},
		{0, 0, 0, 0, 0, 0, 0},
		{1, 1, 1, 1, 1, 0, 1},
	}
	expect := 4
	runSample(t, n, m, G, expect)
}

func TestSample4(t *testing.T) {
	n, m := 3, 5
	G := [][]int{
		{1, 0, 1, 0, 0},
		{1, 1, 1, 1, 0},
		{0, 0, 1, 0, 0},
	}
	expect := 4
	runSample(t, n, m, G, expect)
}
