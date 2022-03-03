package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, E [][]int, C []int, X []int, expect []int) {
	res := solve(n, E, C, X)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 6
	C := []int{1, 2, 3, 2, 6, 3}
	X := []int{2, 1, 8, 1, 2, 1}
	E := [][]int{
		{1, 2},
		{1, 3},
		{2, 4},
		{2, 5},
		{3, 6},
	}
	expect := []int{2, 1, 0, 1, 0, 1}
	runSample(t, n, E, C, X, expect)
}
