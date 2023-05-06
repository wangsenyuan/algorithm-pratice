package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, X []int, W []int, Q [][]int, expect []int64) {
	res := solve(X, W, Q)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	X := []int{-2, 0, 1, 9, 12}
	W := []int{2, 10, 1, 2, 7}
	Q := [][]int{
		{1, 3},
		{2, 3},
		{1, 5},
		{3, 5},
		{2, 4},
	}
	expect := []int64{9, 11, 9, 24, 11}
	runSample(t, X, W, Q, expect)
}
