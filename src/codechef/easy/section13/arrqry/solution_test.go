package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, N, M, Q int, A []int, B []int, queries [][]int, expect []int64) {
	res := solve(N, M, Q, A, B, queries)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %d %d %d %v %v %v, expect %v, but got %v", N, M, Q, A, B, queries, expect, res)
	}
}

func TestSample1(t *testing.T) {
	N, M, Q := 3, 4, 6
	A := []int{2, -1, 5}
	B := []int{3, 3, 2, 4}
	quries := [][]int{
		{3},
		{1, 2, 3, -2},
		{3},
		{1, 1, 3, 1},
		{2, 2, 4, 2},
		{3},
	}
	expect := []int64{72, 24, 90}
	runSample(t, N, M, Q, A, B, quries, expect)
}
