package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, E [][]int, X [][]int, expect []int) {
	res := solve(n, E, X)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 5
	E := [][]int{
		{3, 2},
		{4, 2},
		{4, 1},
		{5, 3},
	}
	X := [][]int{
		{1, 1},
		{4, 5},
		{1, 5},
		{5, 1},
		{3, 3},
	}
	expect := []int{5, 3, 3, 4, 3}
	runSample(t, n, E, X, expect)
}
