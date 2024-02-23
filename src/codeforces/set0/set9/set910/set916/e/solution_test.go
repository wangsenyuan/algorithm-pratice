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
	a := []int{1, 4, 2, 8, 5, 7}
	edges := [][]int{
		{1, 2},
		{3, 1},
		{4, 3},
		{4, 5},
		{3, 6},
	}
	ops := [][]int{
		{3, 1},
		{2, 4, 6, 3},
		{3, 4},
		{1, 6},
		{2, 2, 4, -5},
		{1, 4},
		{3, 3},
	}
	expect := []int{27, 19, 5}
	runSample(t, a, edges, ops, expect)
}

func TestSample2(t *testing.T) {
	a := []int{4, 3, 5, 6}
	edges := [][]int{
		{1, 2},
		{2, 3},
		{3, 4},
	}
	ops := [][]int{
		{3, 1},
		{1, 3},
		{2, 2, 4, 3},
		{1, 1},
		{2, 2, 4, -3},
		{3, 1},
	}
	expect := []int{18, 21}
	runSample(t, a, edges, ops, expect)
}
