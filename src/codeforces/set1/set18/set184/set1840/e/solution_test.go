package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, s1 string, s2 string, x int, queries [][]int, expect []bool) {
	res := solve(x, s1, s2, queries)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s1 := "codeforces"
	s2 := "codeblocks"
	x := 5
	queries := [][]int{
		{3},
		{1, 5},
		{1, 6},
		{1, 7},
		{1, 9},
		{3},
		{3},
	}
	expect := []bool{false, true, false}

	runSample(t, s1, s2, x, queries, expect)
}

func TestSample2(t *testing.T) {
	s1 := "cool"
	s2 := "club"
	x := 2
	queries := [][]int{
		{2, 1, 2, 2, 3},
		{2, 2, 2, 2, 4},
		{1, 2},
		{3},
		{3},
	}
	expect := []bool{true, false}

	runSample(t, s1, s2, x, queries, expect)
}
