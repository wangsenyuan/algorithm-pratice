package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, N, M int, A []int, queries [][]int, expect []int) {
	res := solve(N, M, A, queries)
	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %d %d %v %v, expect %v, but got %v", N, M, A, queries, expect, res)
	}
}

func TestSample1(t *testing.T) {
	N, M := 3, 6
	A := []int{
		encode([]string{"0110", "0000", "0000", "0001"}),
		encode([]string{"0000", "0000", "0000", "0000"}),
		encode([]string{"1000", "1000", "0000", "0000"}),
	}
	queries := [][]int{
		{1, 1, 1},
		{1, 2, 2},
		{1, 3, 3},
		{1, 1, 3},
		{2, 2, encode([]string{"0001", "0010", "0100", "1000"})},
		{1, 1, 3},
	}
	expect := []int{1, 0, 1, 1, 1}

	runSample(t, N, M, A, queries, expect)
}
