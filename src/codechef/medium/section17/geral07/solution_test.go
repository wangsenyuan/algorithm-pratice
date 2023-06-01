package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, N int, M int, E [][]int, Q [][]int, expect []int) {
	res := solve(N, M, E, Q)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m := 3, 5
	E := [][]int{
		{1, 3},
		{1, 2},
		{2, 1},
		{3, 2},
		{2, 2},
	}
	Q := [][]int{
		{2, 3},
		{1, 5},
		{5, 5},
		{1, 2},
	}
	expect := []int{2, 1, 3, 1}
	runSample(t, n, m, E, Q, expect)
}
