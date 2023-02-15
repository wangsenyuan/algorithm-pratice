package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, Q [][]int, expect []bool) {
	res := solve(n, Q)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 8
	Q := [][]int{
		{1, 2, 4},
		{3, 6, 2, 7, 2},
		{1, 3, 2},
		{3, 6, 2, 7, 2},
		{1, 4, 3},
		{3, 2, 6, 4, 8},
		{2, 4, 3},
		{3, 2, 6, 4, 8},
		{1, 4, 8},
		{3, 2, 6, 4, 8},
	}
	expect := []bool{false, true, true, false, true}
	runSample(t, n, Q, expect)
}
