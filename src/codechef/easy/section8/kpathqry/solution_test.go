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
		{1, 3},
		{2, 4},
		{2, 5},
		{3, 6},
	}
	Q := [][]int{
		{6, 2, 4},
		{2, 3, 4, 5},
		{4, 5, 1},
	}
	expect := []bool{true, false, false}

	runSample(t, n, E, Q, expect)
}

func TestSample2(t *testing.T) {
	n := 5
	E := [][]int{
		{1, 2},
		{2, 3},
		{1, 4},
		{4, 5},
	}
	Q := [][]int{
		{2, 3, 4, 5},
	}
	expect := []bool{true}

	runSample(t, n, E, Q, expect)
}
