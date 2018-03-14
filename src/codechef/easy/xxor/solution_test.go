package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, A []int, queries [][]int, expect []int) {
	res := solve(n, A, queries)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("sample %d %v %v, expect %v, but got %v", n, A, queries, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 5
	A := []int{20, 11, 18, 2, 13}
	queries := [][]int{
		{1, 3},
		{3, 5},
		{2, 4},
	}

	expect := []int{2147483629, 2147483645, 2147483645}

	runSample(t, n, A, queries, expect)
}
