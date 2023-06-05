package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, A []int, Q [][]int, expect []int) {
	res := solve(A, Q)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{5, 3, 2, 6, 0}
	Q := [][]int{
		{2, 1, 5},
		{1, 2, 7},
		{2, 1, 5},
		{1, 4, 2},
		{2, 1, 5},
		{1, 2, 5},
		{2, 1, 2},
	}
	expect := []int{4, -1, 5, 0}
	runSample(t, A, Q, expect)
}
