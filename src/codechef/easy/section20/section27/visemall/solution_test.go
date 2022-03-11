package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, S string, Q [][]int, expect []bool) {
	res := solve(n, S, Q)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 5
	S := "11011"
	Q := [][]int{
		{2, 1, 5, 3},
		{2, 3, 5, 3},
		{2, 3, 5, 4},
		{1, 4, 5},
		{2, 1, 5, 2},
		{1, 2, 3},
		{2, 3, 5, 4},
	}
	expect := []bool{true, false, true, false, true}
	runSample(t, n, S, Q, expect)
}
