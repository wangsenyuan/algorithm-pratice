package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n, m int, grid [][]int, expect []string) {
	res := solve(n, m, grid)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %d %d %v, expect %v, but got %v", n, m, grid, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m := 3, 4
	grid := [][]int{
		{1, 3, 7, 10},
		{9, 2, 4, 11},
		{8, 12, 5, 6},
	}

	expect := []string{
		"1111",
		"1001",
		"0100",
	}
	runSample(t, n, m, grid, expect)
}
