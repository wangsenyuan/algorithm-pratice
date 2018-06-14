package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, edges [][]int, K int, queries [][]int, expect []bool) {
	res := solve(n, edges, K, queries)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %d %v %d %v, expect %v, but got %v", n, edges, K, queries, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 10
	edges := [][]int{
		{3, 1},
		{5, 7},
		{4, 6},
		{10, 4},
		{8, 9},
		{2, 1},
		{3, 4},
		{8, 5},
		{5, 3},
	}
	K := 4
	queries := [][]int{
		{8, 10, 3},
		{8, 7, 9, 6},
		{2, 1, 6},
		{6, 2, 5},
	}
	expect := []bool{true, false, true, false}

	runSample(t, n, edges, K, queries, expect)
}

func TestSample2(t *testing.T) {
	n := 6
	edges := [][]int{
		{1, 2},
		{2, 3},
		{3, 4},
		{3, 5},
		{5, 6},
	}
	K := 1
	queries := [][]int{
		{1, 2, 3, 4, 6},
	}
	expect := []bool{false}

	runSample(t, n, edges, K, queries, expect)
}

func TestSample3(t *testing.T) {
	n := 6
	edges := [][]int{
		{1, 2},
		{1, 3},
		{2, 4},
		{3, 5},
		{5, 6},
	}
	K := 1
	queries := [][]int{
		{2, 4, 6},
	}
	expect := []bool{true}

	runSample(t, n, edges, K, queries, expect)
}
