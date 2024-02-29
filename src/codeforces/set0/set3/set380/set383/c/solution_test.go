package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, a []int, edges [][]int, ops [][]int, expect []int) {
	res := solve(a, edges, ops)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{1, 2, 1, 1, 2}
	edges := [][]int{
		{1, 2},
		{1, 3},
		{2, 4},
		{2, 5},
	}
	ops := [][]int{
		{1, 2, 3},
		{1, 1, 2},
		{2, 1},
		{2, 2},
		{2, 4},
	}
	expect := []int{3, 3, 0}
	runSample(t, a, edges, ops, expect)
}
