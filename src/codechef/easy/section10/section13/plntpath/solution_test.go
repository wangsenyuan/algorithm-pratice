package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, K, S int, A []int, C [][]int, expect []int64) {
	res := solve(K, S, A, C)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	K, S := 3, 4
	A := []int{1, 3, 1, 2, 2}
	C := [][]int{
		{0, 5, -1},
		{-1, 4, 0},
		{0, 3, -1},
		{4, 0, -1},
		{4, 0, 3},
	}
	expect := []int64{4, 3, 4, 0, 0}

	runSample(t, K, S, A, C, expect)
}
