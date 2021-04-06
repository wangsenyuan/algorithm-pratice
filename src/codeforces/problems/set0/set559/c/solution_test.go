package main

import "testing"

func runSample(t *testing.T, h, w int, P [][]int, expect int) {
	res := solve(h, w, P)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	h, w := 3, 4
	P := [][]int{
		{2, 2},
		{2, 3},
	}
	expect := 2
	runSample(t, h, w, P, expect)
}

func TestSample2(t *testing.T) {
	h, w := 100, 100
	P := [][]int{
		{15, 16},
		{16, 15},
		{99, 88},
	}
	expect := 545732279
	runSample(t, h, w, P, expect)
}
