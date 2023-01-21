package main

import "testing"

func runSample(t *testing.T, n, k int, A string, C [][]int, expect string) {
	res := solve(n, k, A, C)

	if res != expect {
		t.Errorf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, k := 4, 3
	A := "abab"
	C := [][]int{
		{2, 3},
		{0, 0},
		{0, 4},
		{0, 0},
	}
	expect := "baaaab"
	runSample(t, n, k, A, C, expect)
}

func TestSample2(t *testing.T) {
	n, k := 8, 2
	A := "kadracyn"
	C := [][]int{
		{2, 5},
		{3, 4},
		{0, 0},
		{0, 0},
		{6, 8},
		{0, 7},
		{0, 0},
		{0, 0},
	}
	expect := "daarkkcyan"
	runSample(t, n, k, A, C, expect)
}

func TestSample3(t *testing.T) {
	n, k := 8, 3
	A := "kdaracyn"
	C := [][]int{
		{2, 5},
		{0, 3},
		{0, 4},
		{0, 0},
		{6, 8},
		{0, 7},
		{0, 0},
		{0, 0},
	}
	expect := "darkcyan"
	runSample(t, n, k, A, C, expect)
}
