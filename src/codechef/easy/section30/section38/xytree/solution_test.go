package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, E [][]int, Q [][]int, expect []int) {
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
		{2, 5},
	}
	Q := [][]int{
		{1, 1},
		{1, 3},
		{2},
		{1, 5},
		{2},
	}
	expect := []int{3, 2}
	runSample(t, n, E, Q, expect)
}
