package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, s string, queries [][]int, expect []bool) {
	res := solve(s, queries)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "aaaaa"
	queries := [][]int{
		{1, 1},
		{2, 4},
		{5, 5},
	}
	expect := []bool{true, false, true}

	runSample(t, s, queries, expect)
}

func TestSample2(t *testing.T) {
	s := "aabbbbbbc"
	queries := [][]int{
		{1, 2},
		{2, 4},
		{2, 2},
		{1, 9},
		{5, 7},
		{3, 5},
	}
	expect := []bool{false, true, true, true, false, false}

	runSample(t, s, queries, expect)
}
