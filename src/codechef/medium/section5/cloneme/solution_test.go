package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, A []int, queries [][]int, expect []bool) {
	res := solve(n, len(queries), A, queries)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %d %v %v, expect %v, but got %v", n, A, queries, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 6
	A := []int{1, 3, 4, 2, 3, 4}
	queries := [][]int{
		{1, 3, 4, 6},
		{1, 2, 5, 6},
		{3, 5, 2, 4},
	}
	expect := []bool{true, false, true}
	runSample(t, n, A, queries, expect)
}
