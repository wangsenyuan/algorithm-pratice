package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n, R int, E [][]int, Q [][]int, expect []int64) {
	res := solve(n, R, E, Q)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, R := 3, 1
	E := [][]int{
		{1, 2, 5},
		{1, 3, 4},
	}
	Q := [][]int{
		{2, 3},
		{1, 3},
	}
	expect := []int64{9, 4}
	runSample(t, n, R, E, Q, expect)
}
