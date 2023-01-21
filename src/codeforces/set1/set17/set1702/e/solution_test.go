package main

import "testing"

func runSample(t *testing.T, dominoes [][]int, expect bool) {
	res := solve(dominoes)

	if res != expect {
		t.Errorf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	dominoes := [][]int{
		{1, 2},
		{4, 3},
		{2, 1},
		{3, 4},
	}
	expect := true

	runSample(t, dominoes, expect)
}

func TestSample2(t *testing.T) {
	dominoes := [][]int{
		{1, 2},
		{4, 5},
		{1, 3},
		{4, 6},
		{2, 3},
		{5, 6},
	}
	expect := false

	runSample(t, dominoes, expect)
}

func TestSample3(t *testing.T) {
	dominoes := [][]int{
		{1, 1},
		{2, 2},
	}
	expect := false

	runSample(t, dominoes, expect)
}

func TestSample4(t *testing.T) {
	dominoes := [][]int{
		{1, 2},
		{2, 1},
	}
	expect := true

	runSample(t, dominoes, expect)
}

func TestSample5(t *testing.T) {
	dominoes := [][]int{
		{2, 1},
		{1, 2},
		{4, 3},
		{4, 3},
		{5, 6},
		{5, 7},
		{8, 6},
		{7, 8},
	}
	expect := true

	runSample(t, dominoes, expect)
}

func TestSample6(t *testing.T) {
	dominoes := [][]int{
		{1, 2},
		{2, 1},
		{4, 3},
		{5, 3},
		{5, 4},
		{6, 7},
		{8, 6},
		{7, 8},
	}
	expect := false

	runSample(t, dominoes, expect)
}
