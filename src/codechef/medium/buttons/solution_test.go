package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, src [][]int, dst [][]int, expect bool, row []int, col []int) {
	res, a, b := solve(n, src, dst)

	if res != expect || !reflect.DeepEqual(row, a) || !reflect.DeepEqual(col, b) {
		t.Errorf("Sample %d %v %v, expect %t %v %v, but got %t %v %v", n, src, dst, expect, row, col, res, a, b)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	src := [][]int{
		{0, 0, 0},
		{1, 1, 0},
		{1, 1, 0},
	}
	dst := [][]int{
		{1, 1, 0},
		{1, 1, 1},
		{1, 1, 1},
	}
	row := []int{0}
	col := []int{2}
	runSample(t, n, src, dst, true, row, col)
}

func TestSample2(t *testing.T) {
	n := 5
	src := [][]int{
		{0, 1, 1, 1, 1},
		{1, 0, 0, 0, 0},
		{1, 0, 0, 0, 0},
		{1, 0, 0, 0, 0},
		{0, 1, 1, 1, 1},
	}
	dst := [][]int{
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
	}
	row := []int{0, 4}
	col := []int{0}
	runSample(t, n, src, dst, true, row, col)
}
