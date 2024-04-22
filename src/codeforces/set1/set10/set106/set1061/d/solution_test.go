package main

import (
	"testing"
)

func runSample(t *testing.T, x int, y int, shows [][]int, expect int) {
	res := solve(x, y, shows)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	x, y := 4, 3
	shows := [][]int{
		{1, 2},
		{4, 10},
		{2, 4},
		{10, 11},
		{5, 9},
	}
	expect := 60
	runSample(t, x, y, shows, expect)
}

func TestSample2(t *testing.T) {
	x, y := 3, 2
	shows := [][]int{
		{8, 20},
		{6, 22},
		{4, 15},
		{20, 28},
		{17, 25},
		{20, 27},
	}
	expect := 142
	runSample(t, x, y, shows, expect)
}

func TestSample3(t *testing.T) {
	x, y := 1000000000, 2
	shows := [][]int{
		{1, 2},
		{2, 3},
	}
	expect := 999999997
	runSample(t, x, y, shows, expect)
}
