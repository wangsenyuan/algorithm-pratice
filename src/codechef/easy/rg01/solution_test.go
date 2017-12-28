package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, V int, edges [][]int, expect [][]int) {
	grid := solve(V, len(edges), edges)
	if !reflect.DeepEqual(expect, grid) {
		t.Errorf("sample %d, %v should give answer %v, but got %v", V, edges, expect, grid)
	}
}

func TestSample(t *testing.T) {
	V := 3
	edges := [][]int{
		{0, 1, 1},
		{1, 2, 2},
		{0, 2, 3},
	}
	expect := [][]int{
		{0, 2, 3},
		{2, 0, 2},
		{3, 2, 0},
	}
	runSample(t, V, edges, expect)
}
