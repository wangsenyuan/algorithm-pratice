package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, E [][]int, A []int, Q [][]int, expect []int) {
	res := solve(n, E, A, Q)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 6
	A := []int{1, 2, 3, 4, 5, 6}
	E := [][]int{
		{1, 2},
		{1, 3},
		{2, 4},
		{2, 5},
		{3, 6},
	}
	Q := [][]int{
		{3, 3},
		{3, 2},
		{1, 5, 2},
		{3, 2},
	}
	expect := []int{1, 0, 1}
	runSample(t, n, E, A, Q, expect)
}
