package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, A []int64, K []int, Q [][]int64, expect []int64) {
	res := make([]int64, len(K))
	solver := NewSolver(A)

	for i := 0; i < len(K); i++ {
		res[i] = solver.Query(K[i], Q[i])
	}

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int64{15, 5, 26, 4, 3, 13, 5, 5, 27, 11}
	K := []int{3, 1}
	Q := [][]int64{
		{2, 3, 1, 4, 7, 1, 5, 1, 1, 1},
		{0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
	}
	expect := []int64{33, 5}
	runSample(t, A, K, Q, expect)
}
