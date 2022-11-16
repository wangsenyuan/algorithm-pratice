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
	A := []int{2, 4, 6, 8, 9}
	Q := [][]int{
		{1, 1, 3},
		{2, 2, 56},
		{1, 1, 3},
	}
	expect := []int64{6, 2}
	runSample(t, A, Q, expect)
}

func TestSample2(t *testing.T) {
	A := []int{20, 8, 50, 16, 22, 23, 29}
	Q := [][]int{
		{1, 1, 6},
		{2, 2, 18},
		{1, 1, 4},
		{2, 6, 10},
		{2, 5, 11},
		{1, 5, 6},
	}
	expect := []int64{3, 3, 5}
	runSample(t, A, Q, expect)
}
