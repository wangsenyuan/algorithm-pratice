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
	A := []int{1, 10, 3, 6, 5}
	Q := [][]int{
		{1, 1, 3, 5},
		{2, 3},
		{1, 4, 5, 7},
		{2, 5},
	}
	expect := []int64{52, 69}
	runSample(t, n, A, Q, expect)
}
