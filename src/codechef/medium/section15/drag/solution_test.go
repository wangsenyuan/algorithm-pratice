package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, A []int64, Q [][]int, expect []int) {
	res := solve(A, Q)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int64{2, 3, 4, 4}
	Q := [][]int{
		{1, 3},
		{1, 4},
		{3, 4},
	}
	expect := []int{5, 7, 2}
	runSample(t, A, Q, expect)
}
