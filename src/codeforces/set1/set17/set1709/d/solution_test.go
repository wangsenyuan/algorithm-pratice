package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, m int, A []int, Q [][]int, expect []bool) {
	res := solve(n, m, A, Q)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m := 11, 10
	A := []int{9, 0, 0, 10, 3, 4, 8, 11, 10, 8}
	Q := [][]int{
		{1, 2, 1, 3, 1},
		{1, 2, 1, 3, 2},
		{4, 3, 4, 5, 2},
		{5, 3, 11, 5, 3},
		{5, 3, 11, 5, 2},
		{11, 9, 9, 10, 1},
	}
	expect := []bool{true, false, false, false, true, true}

	runSample(t, n, m, A, Q, expect)
}
