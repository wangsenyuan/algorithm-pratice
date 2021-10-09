package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, A, B []int, Q [][]int, expect []int64) {
	res := solve(A, B, Q)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{4, 6, 5, 2, 7}
	B := []int{3, 2, 4, 6, 3}
	Q := [][]int{
		{2, 4, 2},
		{1, 1, 1},
	}
	expect := []int64{51, 12}
	runSample(t, A, B, Q, expect)
}
