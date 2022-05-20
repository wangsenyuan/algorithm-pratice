package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, A []int, Q [][]int, expect []int) {
	res := solve(n, A, Q)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 6
	A := []int{5, 1, 3, 2, 4, 6}
	Q := [][]int{
		{1, 2},
		{2, 3},
		{1, 5},
		{3, 6},
	}
	expect := []int{2, 1, 5, 3}
	runSample(t, n, A, Q, expect)
}
