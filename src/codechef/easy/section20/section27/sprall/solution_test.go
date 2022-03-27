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
		{1, 3},
		{1, 4},
	}
	Q := [][]int{
		{1, 2, 2, 3},
		{1, 3, 2, 3},
		{4, 1, 1, 4},
		{2, 1, 4, 2},
	}
	expect := []int64{6, 5, 5, 6, 5}
	runSample(t, n, E, Q, expect)
}
