package main

import (
	"testing"
	"reflect"
)

func runSample(t *testing.T, n int, S string, queries [][]int, expect []int) {
	res := solve(n, []byte(S), queries)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("sample %d %s %v, expect %v, but got %v", n, S, queries, expect, res)
	}
}

func TestSample(t *testing.T) {
	n := 5
	S := "01245"
	queries := [][]int{
		{2, 1, 3},
		{1, 4, 5},
		{2, 3, 5},
	}
	expect := []int{3, 1}
	runSample(t, n, S, queries, expect)
}
