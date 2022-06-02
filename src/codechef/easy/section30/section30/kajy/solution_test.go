package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, A, B, C []int, Q [][]int, expect []int64) {
	res := solve(A, B, C, Q)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{10, 1, 6, 10, 2, 2, 5}
	B := []int{4, 9, 4, 4, 3, 6, 6}
	C := []int{5, 1, 10, 10, 7, 2, 6}
	Q := [][]int{
		{2, 2, 9},
		{3, 4, 7},
		{2, 6, 4},
	}
	expect := []int64{54, 51, 51}
	runSample(t, A, B, C, Q, expect)
}
