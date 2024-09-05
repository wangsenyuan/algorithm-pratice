package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, a []int, edges [][]int, expect []int) {
	res := solve(n, a, edges)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 9
	a := []int{0, 1, 1, 1, 0, 0, 0, 0, 1}
	edges := [][]int{
		{1, 2},
		{1, 3},
		{3, 4},
		{3, 5},
		{2, 6},
		{4, 7},
		{6, 8},
		{5, 9},
	}
	expect := []int{2, 2, 2, 2, 2, 1, 1, 0, 2}
	runSample(t, n, a, edges, expect)
}

func TestSample2(t *testing.T) {
	n := 4
	a := []int{0, 0, 1, 0}
	edges := [][]int{
		{1, 2},
		{1, 3},
		{1, 4},
	}
	expect := []int{0, -1, 1, -1}
	runSample(t, n, a, edges, expect)
}
