package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, edges [][]int, special []int, expect []int) {
	res := solve(n, edges, special)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 2
	special := []int{2, 1}
	edges := [][]int{
		{1, 2, 3},
		{1, 2, 2},
		{2, 2, 1},
	}
	expect := []int{2, 2}
	runSample(t, n, edges, special, expect)
}

func TestSample2(t *testing.T) {
	n := 4
	special := []int{1, 2, 3}
	edges := [][]int{
		{1, 2, 5},
		{4, 2, 1},
		{2, 3, 2},
		{1, 4, 4},
		{1, 3, 3},
	}
	expect := []int{3, 3, 3}
	runSample(t, n, edges, special, expect)
}
