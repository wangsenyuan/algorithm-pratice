package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, itvs [][]int, q int, queries [][]int, expect []int) {
	res := solve(n, itvs, q, queries)
	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %d %v %d %v, expect %v, but got %v", n, itvs, q, queries, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 5
	itvs := [][]int{
		{4, 5},
		{3, 5},
		{2, 4},
		{1, 3},
		{5, 5},
	}
	q := 2
	queries := [][]int{
		{1, 2, 3, 4},
		{4},
	}
	expect := []int{3, 3}
	runSample(t, n, itvs, q, queries, expect)
}

func TestSample2(t *testing.T) {
	n := 5
	itvs := [][]int{
		{4, 5},
		{3, 5},
		{2, 4},
		{2, 3},
		{5, 5},
	}
	q := 2
	queries := [][]int{
		{2, 5},
		{1, 2, 5},
	}
	expect := []int{5, 5}
	runSample(t, n, itvs, q, queries, expect)
}
