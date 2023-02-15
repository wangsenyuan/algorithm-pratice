package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, A []int, Q [][]int, expect []int64) {
	res := solve(A, Q)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{1, 2, 3, 4, 5}
	Q := [][]int{
		{1, 1, 5},
		{2, 10},
		{1, 5, 11},
		{1, 4, 1},
		{2, 1},
	}
	expect := []int64{19, 50, 51, 42, 5}
	runSample(t, A, Q, expect)
}
