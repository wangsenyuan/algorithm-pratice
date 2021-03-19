package main

import "testing"

func runSample(t *testing.T, red, blue [][]int, expect int) {
	res := solve(red, blue)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	red := [][]int{
		{0, 0},
		{10, 10},
		{0, 10},
		{10, 0},
	}
	blue := [][]int{
		{12, 11},
		{2, 1},
		{12, 1},
		{2, 11},
	}
	expect := 2
	runSample(t, red, blue, expect)
}
