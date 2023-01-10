package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, A []int, ops [][]int, expect []int64) {
	res := solve(A, ops)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{1, 2, 3, 4}
	ops := [][]int{
		{3, 0},
		{3, 0, -1},
		{0, 1},
		{2, 1},
	}
	expect := []int64{1, 0, 0}
	runSample(t, A, ops, expect)
}
