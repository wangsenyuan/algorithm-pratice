package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, q int, C []int, E [][]int, Q [][]int, expect []int) {
	res := solve(n, q, C, E, Q)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, q := 6, 7
	C := []int{1, 2, 3, 4, 5, 6}
	E := [][]int{
		{1, 2},
		{1, 3},
		{2, 4},
		{2, 5},
		{3, 6},
	}
	Q := [][]int{
		{1, 5, 6},
		{2, 5, 6},
		{1, 1, 5},
		{1, 5, 6},
		{2, 1, 6},
		{1, 5, 6},
		{1, 2, 6},
	}
	expect := []int{5, 3, 4, 3, 3}
	runSample(t, n, q, C, E, Q, expect)
}
