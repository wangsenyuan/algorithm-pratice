package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n, m int, bad [][]int, Q [][]int, expect []int) {
	res := solve(n, m, bad, Q)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m := 4, 5
	bad := [][]int{
		{2, 1},
		{3, 2},
		{1, 3},
		{2, 4},
		{1, 4},
	}
	Q := [][]int{
		{1, 1},
		{3, 4},
		{2, 3},
		{1, 4},
	}
	expect := []int{2, 6, 3, 6}
	runSample(t, n, m, bad, Q, expect)
}
