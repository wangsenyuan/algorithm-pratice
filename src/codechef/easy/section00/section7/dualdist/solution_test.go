package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, E [][]int, Q [][]int, expect []int64) {
	res := solve(n, E, Q)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 4
	E := [][]int{
		{1, 2},
		{2, 3},
		{3, 4},
	}
	Q := [][]int{
		{1, 2},
		{3, 1},
	}

	expect := []int64{3, 2}
	runSample(t, n, E, Q, expect)
}

func TestSample2(t *testing.T) {
	n := 7
	E := [][]int{
		{1, 2},
		{1, 3},
		{2, 4},
		{2, 5},
		{3, 6},
		{3, 7},
	}
	Q := [][]int{
		{2, 6},
	}

	expect := []int64{6}
	runSample(t, n, E, Q, expect)
}
