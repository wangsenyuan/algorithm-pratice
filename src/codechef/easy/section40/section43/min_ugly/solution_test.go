package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, E [][]int, Q [][]int, expect []int) {
	res := solve(n, E, Q)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 7
	E := [][]int{
		{1, 2},
		{1, 3},
		{2, 4},
		{2, 5},
		{3, 6},
		{3, 7},
	}
	Q := [][]int{
		{4},
		{4, 7, 5, 6},
	}
	expect := []int{0, 2}
	runSample(t, n, E, Q, expect)
}
