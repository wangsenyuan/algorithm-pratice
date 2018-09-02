package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, N int, A []int, M int, queries []int, expect []int) {
	res := solve(N, A, M, queries)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %d %v %d %v, expect %v, but got %v", N, A, M, queries, expect, res)
	}
}

func TestSample1(t *testing.T) {
	N, M := 5, 5
	A := []int{1, 4, 3, 5, 2}
	queries := []int{2, 2, 4, 4, 5, 5, 3, 7, 4, 8}
	expect := []int{1, 3, 4, 5, 2, 2}
	runSample(t, N, A, M, queries, expect)
}
