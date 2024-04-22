package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, s string, queries [][]int, expect []int) {
	res := solve(s, queries)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample %s %v, expect %v, but got %v", s, queries, expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "1011"
	queries := [][]int{
		{1, 4},
		{3, 4},
	}

	// 1110
	// 1, 221
	// 1 + 2, 43
	// 1 + 2 + 4, 7
	// 1 + 2 + 4 + 7

	expect := []int{14, 3}
	runSample(t, s, queries, expect)
}
