package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, p []int, edges [][]int, cmds [][]int, expect []int) {
	res := solve(n, p, edges, cmds)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 5
	p := []int{1, 2, 5, 4, 3}
	edges := [][]int{
		{1, 2},
		{2, 3},
		{1, 3},
		{4, 5},
	}
	cmds := [][]int{
		{1, 1},
		{2, 1},
		{2, 3},
		{1, 1},
		{1, 2},
		{1, 2},
	}
	expect := []int{5, 1, 2, 0}
	runSample(t, n, p, edges, cmds, expect)
}

func TestSample2(t *testing.T) {
	n := 6
	p := []int{1, 2, 6, 3, 4, 5}
	edges := [][]int{
		{2, 4},
		{1, 3},
		{1, 6},
		{3, 5},
		{2, 5},
	}
	cmds := [][]int{
		{1, 1},
		{1, 4},
		{1, 1},
		{2, 1},
		{2, 5},
		{2, 2},
		{2, 3},
		{1, 4},
	}
	expect := []int{6, 5, 4, 3}
	runSample(t, n, p, edges, cmds, expect)
}
