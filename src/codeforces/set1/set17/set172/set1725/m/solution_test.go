package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, E [][]int, expect []int64) {
	res := solve(n, E)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 5
	E := [][]int{
		{1, 2, 2},
		{2, 4, 1},
		{4, 1, 4},
		{2, 5, 3},
		{5, 4, 1},
		{5, 2, 4},
		{2, 1, 1},
	}
	expect := []int64{1, -1, 3, 4}
	runSample(t, n, E, expect)
}
