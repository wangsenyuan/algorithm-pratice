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
	A := []int{2, 1, 3, 4}
	Q := [][]int{
		{1, 1},
		{1, 2},
		{2, 1},
		{2, 2},
		{3, 2},
	}
	expect := []int{2, 1, 2, 1, 0}
	runSample(t, A, Q, expect)
}
