package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, grid []string, queries [][]int, expect []int) {
	res := solve(grid, queries)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	grid := []string{
		"000",
		"111",
		"000",
	}
	queries := [][]int{
		{1, 1, 1},
		{2, 2, 2},
		{3, 3, 3},
	}
	expect := []int{1, 1, 1}
	runSample(t, grid, queries, expect)
}

func TestSample2(t *testing.T) {
	grid := []string{
		"01",
		"10",
		"01",
		"10",
		"01",
	}
	queries := [][]int{
		{1, 1, 4},
		{5, 1, 4},
	}
	expect := []int{0, 0}
	runSample(t, grid, queries, expect)
}

func TestSample3(t *testing.T) {
	grid := []string{
		"01011",
		"10110",
		"01101",
		"11010",
		"10101",
	}
	queries := [][]int{
		{1, 1, 4},
		{1, 2, 3},
		{5, 5, 3},
	}
	expect := []int{1, 0, 1}
	runSample(t, grid, queries, expect)
}

func TestSample4(t *testing.T) {
	grid := []string{
		"0",
	}
	queries := [][]int{
		{1, 1, 1},
		{1, 1, 2},
		{1, 1, 3},
	}
	expect := []int{0, 0, 0}
	runSample(t, grid, queries, expect)
}
