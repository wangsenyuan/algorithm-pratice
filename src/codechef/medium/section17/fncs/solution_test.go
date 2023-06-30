package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, A []int, fns [][]int, Q [][]int, expect []int64) {
	res := solve(A, fns, Q)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{1, 2, 3, 4, 5}
	fns := [][]int{
		{1, 3},
		{2, 5},
		{4, 5},
		{3, 5},
		{1, 2},
	}
	Q := [][]int{
		{2, 1, 4},
		{1, 3, 7},
		{2, 1, 4},
		{2, 3, 5},
	}
	expect := []int64{41, 53, 28}
	runSample(t, A, fns, Q, expect)
}
