package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, ctrls [][]int, fences [][]int, queries [][]int, expect []int) {
	res := solve(ctrls, fences, queries)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	ctrls := [][]int{
		{0, 0},
		{3, 3},
	}
	fences := [][]int{
		{2, 0, 0},
	}
	queries := [][]int{
		{1, 2},
	}
	expect := []int{1}
	runSample(t, ctrls, fences, queries, expect)
}

func TestSample2(t *testing.T) {
	ctrls := [][]int{
		{0, 0},
		{4, 4},
	}
	fences := [][]int{
		{1, 0, 0},
		{2, 0, 0},
		{3, 0, 0},
	}
	queries := [][]int{
		{1, 2},
	}
	expect := []int{3}
	runSample(t, ctrls, fences, queries, expect)
}
