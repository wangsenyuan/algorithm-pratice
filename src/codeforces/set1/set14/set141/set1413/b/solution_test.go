package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, rows [][]int, cols [][]int, expect [][]int) {
	res := solve(len(rows), len(cols), rows, cols)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	rows := [][]int{
		{6, 5, 4},
		{1, 2, 3},
	}
	cols := [][]int{
		{1, 6},
		{2, 5},
		{3, 4},
	}
	expect := [][]int{
		{1, 2, 3},
		{6, 5, 4},
	}
	runSample(t, rows, cols, expect)
}

func TestSample2(t *testing.T) {
	rows := [][]int{
		{2},
		{3},
		{1},
	}
	cols := [][]int{
		{3, 1, 2},
	}
	expect := [][]int{
		{3},
		{1},
		{2},
	}
	runSample(t, rows, cols, expect)
}
