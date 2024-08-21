package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, edges [][]int, expect []int) {
	res := solve(n, edges)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 6
	edges := [][]int{
		{1, 2},
		{1, 3},
		{2, 5},
		{2, 4},
		{5, 1},
		{3, 6},
		{6, 2},
	}
	expect := []int{0, 0, 1, 2, 0, 1}
	runSample(t, n, edges, expect)
}

func TestSample2(t *testing.T) {
	n := 6
	edges := [][]int{
		{1, 2},
		{1, 5},
		{2, 6},
		{6, 1},
		{2, 3},
		{3, 4},
		{4, 2},
		{5, 4},
	}
	expect := []int{0, 0, 2, 1, 1, 0}
	runSample(t, n, edges, expect)
}

func TestSample3(t *testing.T) {
	n := 7
	edges := [][]int{
		{1, 4},
		{4, 2},
		{4, 7},
		{1, 6},
		{7, 5},
		{5, 3},
		{3, 4},
		{3, 1},
		{4, 3},
		{4, 5},
	}
	expect := []int{0, 2, 0, 0, 2, 1, 2}
	runSample(t, n, edges, expect)
}
