package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, mat []string, queries [][]int, expect []bool) {
	res := solve(mat, queries)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	mat := []string{
		"..XXX",
		"...X.",
		"...X.",
		"...X.",
	}
	queries := [][]int{
		{1, 3},
		{3, 3},
		{4, 5},
		{5, 5},
		{1, 5},
	}
	expect := []bool{true, true, false, true, false}
	runSample(t, mat, queries, expect)
}
