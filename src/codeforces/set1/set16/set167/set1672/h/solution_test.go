package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, s string, queries [][]int, expect []int) {
	res := solve(s, queries)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "11011"
	queries := [][]int{
		{2, 4},
		{1, 5},
		{3, 5},
	}
	expect := []int{1, 3, 2}
	runSample(t, s, queries, expect)
}

func TestSample2(t *testing.T) {
	s := "1001110110"
	queries := [][]int{
		{1, 10},
		{2, 5},
		{5, 10},
	}
	expect := []int{4, 2, 3}
	runSample(t, s, queries, expect)
}


