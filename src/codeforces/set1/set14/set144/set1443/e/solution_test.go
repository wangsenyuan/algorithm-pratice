package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, Q [][]int, expect []int64) {
	res := solve(n, Q)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 4
	Q := [][]int{
		{1, 2, 4},
		{2, 3},
		{1, 1, 2},
		{1, 3, 4},
	}
	expect := []int64{9, 4, 6}
	runSample(t, n, Q, expect)
}
