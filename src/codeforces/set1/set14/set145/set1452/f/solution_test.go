package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, cnt []int64, Q [][]int64, expect []int64) {
	res := solve(cnt, Q)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	cnt := []int64{0, 1, 0, 0, 1, 0}
	Q := [][]int64{
		{2, 1, 5},
		{2, 4, 18},
		{1, 1, 0},
		{2, 2, 5},
		{2, 0, 17},
		{1, 0, 3},
		{2, 1, 2},
		{1, 1, 4},
		{1, 4, 0},
		{1, 5, 1},
		{2, 2, 8},
	}
	expect := []int64{4, 16, 4, -1, 0, 1}
	runSample(t, cnt, Q, expect)
}
