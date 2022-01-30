package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, P []int, V []int, Q [][]int, expect []int64) {
	res := solve(n, V, P, Q)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 5
	P := []int{1, 1, 2, 2}
	V := []int{1, 2, 3, 4, 5}
	Q := [][]int{
		{3, 4, 5},
		{1, 4, 3},
		{1, 5, 3},
		{1, 2, 4},
		{2, 4, 10},
		{3, 2, 5},
	}
	expect := []int64{11, 20}
	runSample(t, n, P, V, Q, expect)
}
