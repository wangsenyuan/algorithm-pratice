package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, m int, v int, stairs []int, elevators []int, queries [][]int, expect []int) {
	res := solve(n, m, v, stairs, elevators, queries)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m, v := 5, 6, 3
	stairs := []int{2}
	elevators := []int{5}
	queries := [][]int{
		{1, 1, 5, 6},
		{1, 3, 5, 4},
		{3, 3, 5, 3},
	}
	expect := []int{7, 5, 4}
	runSample(t, n, m, v, stairs, elevators, queries, expect)
}
