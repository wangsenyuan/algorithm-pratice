package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, A []int64, Q [][]int, expect []int64) {
	res := solve(A, Q)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int64{1, 2, 4, 3, 2}
	Q := [][]int{
		{1, 1, 3, 5, 5},
		{2, 1, 2, 3, 5},
	}
	expect := []int64{2, 2}
	runSample(t, A, Q, expect)
}

func TestSample2(t *testing.T) {
	A := []int64{3, 5, 6, 13, 12, 20}
	Q := [][]int{
		{1, 1, 4, 5, 6},
		{3, 2, 3, 4, 6},
	}
	expect := []int64{4, 4}
	runSample(t, A, Q, expect)
}
