package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, E [][]int, Q [][]int, expect []int) {
	res := solve(n, E, Q)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 5
	E := [][]int{
		{1, 2, 2},
		{1, 3, 1},
		{1, 4, 4},
		{2, 4, 1},
		{3, 4, 2},
		{4, 5, 0},
	}
	Q := [][]int{
		{2, 4},
		{3, 4},
		{4, 5},
	}
	expect := []int{2, 1, 3}
	runSample(t, n, E, Q, expect)
}
