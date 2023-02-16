package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, A []int, Q [][]int, expect []int) {
	res := solve(A, Q)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{2, 1, 1, 2, 1, 1, 3, 2, 1, 1, 3, 3}
	Q := [][]int{
		{1, 2, 10, 3},
		{1, 2, 11, 3},
		{2, 7, 2},
		{1, 3, 9, 2},
		{1, 1, 12, 1},
		{1, 1, 12, 4},
		{2, 12, 4},
		{1, 1, 12, 4},
		{2, 1, 5},
		{1, 3, 12, 2},
		{1, 1, 4, 3},
	}
	expect := []int{5, 4, 1, 0, -1, 5, 0, 1}
	runSample(t, A, Q, expect)
}
