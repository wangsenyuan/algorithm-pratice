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
	s := "aaab"
	queries := [][]int{
		{1, 4},
		{1, 3},
		{3, 4},
		{2, 4},
	}
	expect := []int{9, 0, 2, 5}
	runSample(t, s, queries, expect)
}

func TestSample2(t *testing.T) {
	s := "abc"
	queries := [][]int{
		{1, 3},
		{1, 2},
	}
	expect := []int{5, 2}
	runSample(t, s, queries, expect)
}

func TestSample3(t *testing.T) {
	s := "pqpcc"
	queries := [][]int{
		{1, 5},
		{4, 5},
		{1, 3},
		{2, 4},
	}
	expect := []int{14, 0, 2, 5}
	runSample(t, s, queries, expect)
}

func TestSample4(t *testing.T) {
	s := "steponnopets"
	queries := [][]int{
		{1, 12},
	}
	expect := []int{65}
	runSample(t, s, queries, expect)
}

func TestSample5(t *testing.T) {
	s := "acaca"
	// 2 + 4
	queries := [][]int{
		{1, 5},
	}
	expect := []int{6}
	runSample(t, s, queries, expect)
}

func TestSample6(t *testing.T) {
	s := "abaccc"
	// 2 + 4
	queries := [][]int{
		{3, 6},
	}
	expect := []int{9}
	runSample(t, s, queries, expect)
}
