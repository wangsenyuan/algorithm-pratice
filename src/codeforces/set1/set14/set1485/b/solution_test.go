package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, q int, k int, A []int, Q [][]int, expect []int) {
	res := solve(n, q, k, A, Q)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, q, k := 4, 2, 5
	A := []int{1, 2, 4, 5}
	Q := [][]int{
		{2, 3},
		{3, 4},
	}
	expect := []int{4, 3}
	runSample(t, n, q, k, A, Q, expect)
}
