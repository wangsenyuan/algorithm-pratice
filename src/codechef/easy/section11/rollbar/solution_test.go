package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, m, n int, x int, y int, grid []string, expect [][]int) {
	res := solve(m, n, x, y, grid)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %d %d %d %d %v, expect %v, but got %v", m, n, x, y, grid, expect, res)
	}
}

func TestSample1(t *testing.T) {
	m, n := 2, 4
	x, y := 1, 1
	grid := []string{
		"1111",
		"0111",
	}
	expect := [][]int{
		{0, -1, -1, 2},
		{-1, -1, -1, 3},
	}
	runSample(t, m, n, x, y, grid, expect)
}

func TestSample2(t *testing.T) {
	m, n := 2, 4
	x, y := 1, 1
	grid := []string{
		"1111",
		"0011",
	}
	expect := [][]int{
		{0, -1, -1, 2},
		{-1, -1, -1, -1},
	}
	runSample(t, m, n, x, y, grid, expect)
}

func TestSample3(t *testing.T) {
	m, n := 4, 2
	x, y := 1, 1
	grid := []string{
		"10",
		"11",
		"11",
		"11",
	}
	expect := [][]int{
		{0, -1},
		{-1, -1},
		{-1, -1},
		{2, 3},
	}
	runSample(t, m, n, x, y, grid, expect)
}