package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, N int, A []int, P []int, Q int, queries [][]int, expect []int) {
	res := solve(N, A, P, Q, queries)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %d %v %v %d %v, expect %v, but got %v", N, A, P, Q, queries, expect, res)
	}
}

func TestSample1(t *testing.T) {
	N := 6
	A := []int{2, 3, 1, 4, 5, 1}
	P := []int{1, 2, 2, 3, 4}
	Q := 5
	queries := [][]int{{2, 1}, {1, 0}, {4, 0}, {5, 2}, {2, 0}}
	expect := []int{2, 1, 3, 3, 0}
	runSample(t, N, A, P, Q, queries, expect)
}
