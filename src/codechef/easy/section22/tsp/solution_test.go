package main

import (
	"testing"
)

func runSample(t *testing.T, n int, C [][]int, D int, expect int) {
	res := solve(n, C, D)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 4
	c := [][]int{
		{0, 0, 2},
		{2, 0},
		{0},
	}
	C := convert(n, c)
	D := 3
	expect := 18
	runSample(t, n, C, D, expect)
}
func TestSample2(t *testing.T) {
	n := 5
	c := [][]int{
		{2, 0, 0, 3},
		{1, 2, 0},
		{3, 3},
		{1},
	}
	C := convert(n, c)
	D := 4
	expect := 38
	runSample(t, n, C, D, expect)
}
