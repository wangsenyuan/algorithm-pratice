package main

import "testing"

func runSample(t *testing.T, N int, grid []string, expect int64) {
	res := solve(N, grid)

	if res != expect {
		t.Errorf("Sample %d %v, expect %d, but got %d", N, grid, expect, res)
	}
}

func TestSample1(t *testing.T) {
	N := 2
	grid := []string{
		"10",
		"01",
	}
	var expect int64 = 3
	runSample(t, N, grid, expect)
}

func TestSample2(t *testing.T) {
	N := 4
	grid := []string{
		"1000",
		"0010",
		"0100",
		"0001",
	}
	var expect int64 = 8
	runSample(t, N, grid, expect)
}
