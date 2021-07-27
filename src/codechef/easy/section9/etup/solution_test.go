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
	A := []int{1, 2, 3, 4, 5, 6}
	Q := [][]int{
		{1, 3},
		{2, 5},
		{1, 6},
	}
	expect := []int64{1, 2, 10}
	runSample(t, A, Q, expect)
}
