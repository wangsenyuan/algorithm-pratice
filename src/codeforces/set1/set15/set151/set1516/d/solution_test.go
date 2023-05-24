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
	A := []int{2, 3, 10, 7, 5, 14}
	Q := [][]int{
		{1, 6},
		{2, 4},
		{3, 5},
	}
	expect := []int{3, 1, 2}
	runSample(t, n, A, Q, expect)
}
