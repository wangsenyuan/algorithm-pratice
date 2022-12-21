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
		{3, 2},
		{4, 2},
		{3, 1},
		{2, 1},
		{2, 2},
	}
	expect := []int64{29, 23, 35, 25, 35}
	runSample(t, A, Q, expect)
}
