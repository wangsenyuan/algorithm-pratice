package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, A []int, Q [][]int, expect []int64) {
	res := solve(n, A, Q)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 5
	A := []int{1, 3, 2, 4, 5}
	Q := [][]int{
		{1, 5},
		{1, 4},
		{2, 4},
		{3, 3},
		{3, 5},
	}
	expect := []int64{4, 3, 2, 0, 5}
	runSample(t, n, A, Q, expect)
}
