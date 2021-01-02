package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, m int, Q [][]int, expect []bool) {
	res := solve(n, m, Q)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m := 4, 5
	Q := [][]int{
		{1, 1, 2, 4, 4},
		{2, 2, 3, 3, 3},
		{2, 1, 1, 4, 5},
	}
	expect := []bool{true, false}
	runSample(t, n, m, Q, expect)
}

func TestSample2(t *testing.T) {
	n, m := 5, 6
	Q := [][]int{
		{1, 2, 2, 5, 5},
		{2, 1, 1, 5, 6},
		{2, 1, 1, 3, 4},
	}
	expect := []bool{true, false}
	runSample(t, n, m, Q, expect)
}
