package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, D [][]int, Q [][]int, expect []int64) {
	res := solve(n, D, Q)
	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 2
	D := [][]int{
		{1, 1, 1, 1},
	}
	Q := [][]int{
		{1, 1, 1, 1},
		{1, 1, 1, 2},
		{1, 1, 2, 1},
		{1, 1, 2, 2},
		{1, 2, 1, 2},
		{1, 2, 2, 1},
		{1, 2, 2, 2},
		{2, 1, 2, 1},
		{2, 1, 2, 2},
		{2, 2, 2, 2},
	}
	expect := []int64{0, 1, 1, 2, 0, 2, 1, 0, 1, 0}
	runSample(t, n, D, Q, expect)
}

func TestSample2(t *testing.T) {
	n := 4
	D := [][]int{
		{1, 1, 1, 1},
		{2, 1, 2, 2},
		{3, 2, 1, 3},
	}
	Q := [][]int{
		{2, 4, 4, 3},
		{4, 4, 3, 3},
		{1, 2, 3, 3},
		{2, 2, 4, 4},
		{1, 4, 2, 3},
	}
	expect := []int64{3, 4, 3, 6, 2}
	runSample(t, n, D, Q, expect)
}

func TestSample3(t *testing.T) {
	n := 5
	D := [][]int{
		{1, 1, 1, 1},
		{2, 1, 1, 2},
		{3, 1, 3, 3},
		{4, 3, 4, 4},
	}
	Q := [][]int{
		{2, 5, 1, 2},
		{1, 1, 5, 2},
		{5, 4, 3, 3},
		{2, 2, 1, 2},
		{2, 5, 3, 5},
	}
	expect := []int64{8, 7, 5, 1, 1}
	runSample(t, n, D, Q, expect)
}
