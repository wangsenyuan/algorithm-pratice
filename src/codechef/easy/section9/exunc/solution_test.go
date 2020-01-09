package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, A []int, Q [][]int, expect []int) {
	res := solve(n, A, Q)
	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %d %v %v, expect %v,  but got %v", n, A, Q, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 5
	A := []int{
		2, 2, 7, 14, 14,
	}
	Q := [][]int{
		{1, 1, 3},
		{1, 2, 6},
		{2, 2},
		{2, 4},
		{2, 5},
	}

	expect := []int{1, 3, 3}

	runSample(t, n, A, Q, expect)
}
