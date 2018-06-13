package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, rules [][]int, m int, queries [][]int, expect []int) {
	res := solve(n, rules, m, queries)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %d %v %d %v, expect %v, but got %v", n, rules, m, queries, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m := 4, 1
	rules := [][]int{
		{22, 1, 200, 2},
		{22, 1, 300, 2},
		{22, 1, 100, 1},
		{22, 1, 400, 2},
	}
	queries := [][]int{{22, 1}}
	expect := []int{400}
	runSample(t, n, rules, m, queries, expect)
}
