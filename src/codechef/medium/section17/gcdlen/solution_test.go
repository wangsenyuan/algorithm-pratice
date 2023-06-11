package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, A []int, Q [][]int, expect []int64) {
	res := solve(A, Q)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{2, 6, 4, 10, 9, 12}
	Q := [][]int{
		{4, 5, 12, 6},
		{2, 5, 2, 4},
		{2, 6, 6, 3},
		{1, 5, 7, 3},
		{1, 6, 4, 0},
	}
	expect := []int64{12, 18, 14, 18, 19, 24}
	runSample(t, A, Q, expect)
}
