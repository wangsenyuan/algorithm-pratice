package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, N int, Q int, E [][]int, QR [][]int, expect []int64) {
	res := solve(N, Q, E, QR)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %d %d %v %v, expect %v, but got %v", N, Q, E, QR, expect, res)
	}
}

func TestSample1(t *testing.T) {
	N, Q := 6, 2
	E := [][]int{
		{1, 2},
		{1, 3},
		{1, 4},
		{2, 5},
		{2, 6},
	}
	queries := [][]int{
		{4, 5}, {1, 3},
	}
	expect := []int64{6, 9}

	runSample(t, N, Q, E, queries, expect)
}
