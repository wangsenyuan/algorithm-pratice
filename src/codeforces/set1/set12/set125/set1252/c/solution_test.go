package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, row []int, col []int, queries [][]int, expect []bool) {
	res := solve(row, col, queries)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	row := []int{6, 2, 7, 8, 3}
	col := []int{3, 4, 8, 5, 1}
	queries := [][]int{
		{2, 2, 1, 3},
		{4, 2, 4, 3},
		{5, 1, 3, 4},
	}

	expect := []bool{true, true, false}

	runSample(t, row, col, queries, expect)
}

func TestSample2(t *testing.T) {
	row := []int{30, 40, 49}
	col := []int{15, 20, 25}
	queries := [][]int{
		{2, 2, 3, 3},
		{1, 2, 2, 2},
	}

	expect := []bool{false, true}

	runSample(t, row, col, queries, expect)
}
