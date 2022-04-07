package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, N int, E [][]int, F [][]int, Q [][]int, expect []int64) {
	res := solve(N, E, F, Q)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 2
	E := [][]int{
		{2, 1, 5},
	}
	F := [][]int{
		{1, 1, 1},
		{5, 4, 1},
	}
	Q := [][]int{
		{2, 2, 2, 1},
		{1, 3, 2, 2},
	}
	expect := []int64{8, 10}
	runSample(t, n, E, F, Q, expect)
}
