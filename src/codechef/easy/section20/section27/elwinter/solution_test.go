package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, E [][]int, Q [][]int, expect []bool) {
	res := solve(n, E, Q)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 6
	E := [][]int{
		{1, 2},
		{2, 3},
		{4, 1},
		{5, 4},
		{6, 1},
		{6, 5},
	}
	Q := [][]int{
		{1, 1},
		{3, 1},
		{1, 5},
		{2, 1},
		{3, 4},
		{3, 3},
	}
	expect := []bool{true, true, false}
	runSample(t, n, E, Q, expect)
}
