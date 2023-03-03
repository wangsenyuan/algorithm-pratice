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
	A := []int{1, 4, 3, 6, 2, 2, 4, 8, 5, 10}
	Q := [][]int{
		{1, 2},
		{1, 3},
		{1, 10},
		{4, 8},
		{2, 5},
	}
	expect := []int64{3, 4, 21, 13, 6}
	runSample(t, A, Q, expect)
}
