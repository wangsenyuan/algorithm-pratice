package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, q int, s string, queries [][]int, expect []bool) {
	res := solve(n, q, s, queries)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %d %d %s %v, expect %v, but got %v", n, q, s, queries, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 10
	q := 2
	s := "helloworld"
	queries := [][]int{
		{1, 3},
		{1, 10},
	}
	expect := []bool{false, true}
	runSample(t, n, q, s, queries, expect)
}
