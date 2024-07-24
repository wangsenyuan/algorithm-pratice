package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, edges [][]int, ord []int, expect []int) {
	res := solve(n, edges, func(m int) []int {
		return ord
	})

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	edges := [][]int{
		{1, 2},
		{2, 3},
	}
	ord := []int{3}
	expect := []int{1, 2, 3, 2, 1}
	runSample(t, n, edges, ord, expect)
}

func TestSample2(t *testing.T) {
	n := 6
	edges := [][]int{
		{1, 2},
		{1, 3},
		{2, 4},
		{4, 5},
		{4, 6},
	}
	ord := []int{5, 6, 3}
	expect := []int{1, 2, 4, 5, 4, 6, 4, 2, 1, 3, 1}
	runSample(t, n, edges, ord, expect)
}
