package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, Q [][]int, expect []int) {
	res := solve(n, Q)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 5
	Q := [][]int{
		{1, 3, 4},
		{1, 2, 1},
		{1, 1, 4},
		{2, 4, 2, 3},
		{2, 4, 2, 1},
		{2, 3, 5, 3},
	}
	expect := []int{4, 1, 3, -1, 4, 5}
	runSample(t, n, Q, expect)
}
