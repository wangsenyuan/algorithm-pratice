package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, A []int, k int, Q [][]int, expect []int) {
	res := solve(A, k, Q)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	k := 1
	A := []int{3, 5, 1, 4, 2}
	Q := [][]int{
		{1, 9},
		{2, 4},
		{3, 6},
		{4, 6},
		{5, 2},
	}
	expect := []int{62, 58, 78, 86, 86}
	runSample(t, A, k, Q, expect)
}
