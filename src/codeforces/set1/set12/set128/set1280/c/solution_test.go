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
		{1, 2, 3},
		{3, 2, 4},
		{2, 4, 3},
		{4, 5, 6},
		{5, 6, 5},
	}
	expect := []int{15, 33}

	runSample(t, n, edges, expect)
}

func TestSample2(t *testing.T) {
	n := 4
	edges := [][]int{
		{1, 2, 1},
		{1, 3, 2},
		{1, 4, 3},
	}
	expect := []int{6, 6}

	runSample(t, n, edges, expect)
}
