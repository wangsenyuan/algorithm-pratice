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
	n := 8
	E := [][]int{
		{1, 2},
		{5, 6},
		{5, 3},
		{4, 3},
		{8, 2},
		{3, 1},
		{7, 5},
	}
	Q := [][]int{
		{3, 5, 2, 3, 2},
		{5, 8, 12, 8, 7},
		{1, 10, 2, 4, 7},
	}
	expect := []int64{6, 80, 20}
	runSample(t, n, E, Q, expect)
}
