package main

import "testing"

func runSample(t *testing.T, P int, G [][]int, expect int) {
	res := solve(len(G), P, G)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	P := 2
	G := [][]int{{0, 1, 2}}
	expect := 1
	runSample(t, P, G, expect)
}

func TestSample2(t *testing.T) {
	P := 13
	G := [][]int{
		{1, 1, 2},
		{3, 1, 3},
		{1, 3, 4},
	}
	expect := 2
	runSample(t, P, G, expect)
}

func TestSample3(t *testing.T) {
	P := 14
	G := [][]int{
		{1, 1, 2},
		{3, 1, 3},
		{1, 3, 4},
	}
	expect := 3
	runSample(t, P, G, expect)
}

func TestSample4(t *testing.T) {
	P := 10
	G := [][]int{
		{1, 3, 4},
		{1, 1, 2},
		{2, 1, 2},
		{2, 3, 4},
	}
	expect := 3
	runSample(t, P, G, expect)
}
