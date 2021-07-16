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
	n := 5
	E := [][]int{
		{1, 2},
		{1, 3},
		{2, 4},
		{4, 5},
	}
	Q := [][]int{
		{2, 2, 3, 5},
		{4, 1, 3},
	}
	expect := []int64{2, 0}
	runSample(t, n, E, Q, expect)
}
