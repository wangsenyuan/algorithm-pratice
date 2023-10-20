package main

import "testing"

func runSample(t *testing.T, grid [][]int, expect int) {
	res := solve(grid)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	grid := [][]int{
		{473507, 133555, 506980, 669045},
		{793254, 26472, 80079, 27306},
		{280175, 374957, 577622, 724612},
		{785695, 64604, 578296, 559883},
	}
	expect := 6575963
	runSample(t, grid, expect)
}

func TestSample2(t *testing.T) {
	grid := [][]int{
		{0, 0, 0, 1000000000},
		{0, 0, 1000000000, 0},
		{0, 1000000000, 0, 0},
		{1000000000, 0, 0, 0},
	}
	expect := 3000000000
	runSample(t, grid, expect)
}

func TestSample3(t *testing.T) {
	grid := [][]int{
		{0, 0, 0, 0, 1000000000},
		{0, 0, 0, 1000000000, 0},
		{0, 0, 1000000000, 0, 0},
		{0, 1000000000, 0, 0, 0},
		{1000000000, 0, 0, 0, 0},
	}
	expect := 4000000000
	runSample(t, grid, expect)
}
