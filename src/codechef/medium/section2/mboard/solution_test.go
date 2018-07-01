package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, N int, Q int, queries []string, expect []int) {
	qs := make([][]byte, Q)
	for i := 0; i < Q; i++ {
		qs[i] = []byte(queries[i])
	}
	res := solve(N, Q, qs)
	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %d %d %v, expect %v, but got %v", N, Q, queries, expect, res)
	}
}

func TestSample1(t *testing.T) {
	N, Q := 3, 6
	queries := []string{
		"RowQuery 1",
		"ColSet 1 1",
		"RowQuery 1",
		"ColQuery 1",
		"RowSet 1 0",
		"ColQuery 1",
	}
	expect := []int{3, 2, 0, 1}
	runSample(t, N, Q, queries, expect)
}
