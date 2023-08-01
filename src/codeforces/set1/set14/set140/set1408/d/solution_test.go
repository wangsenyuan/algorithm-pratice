package main

import "testing"

func runSample(t *testing.T, robbers [][]int, lights [][]int, expect int) {
	res := solve(robbers, lights)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	robs := [][]int{
		{0, 0},
	}
	lights := [][]int{
		{2, 3},
	}
	expect := 3
	runSample(t, robs, lights, expect)
}

func TestSample2(t *testing.T) {
	robs := [][]int{
		{1, 6},
		{6, 1},
	}
	lights := [][]int{
		{10, 1},
		{1, 10},
		{7, 7},
	}
	expect := 4
	runSample(t, robs, lights, expect)
}

func TestSample3(t *testing.T) {
	robs := [][]int{
		{0, 0},
	}
	lights := [][]int{
		{0, 0},
		{0, 0},
	}
	expect := 1
	runSample(t, robs, lights, expect)
}

func TestSample4(t *testing.T) {
	robs := [][]int{
		{0, 8},
		{3, 8},
		{2, 7},
		{0, 10},
		{5, 5},
		{7, 0},
		{3, 5},
	}
	lights := [][]int{
		{6, 6},
		{3, 11},
		{11, 5},
	}
	expect := 6
	runSample(t, robs, lights, expect)
}
