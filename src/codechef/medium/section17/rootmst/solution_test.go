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
	n := 4
	E := [][]int{
		{1, 2, 1},
		{1, 3, 1},
		{1, 4, 1},
		{2, 3, 2},
		{3, 4, 2},
	}
	expect := []int64{5, 4, 3}
	runSample(t, n, E, expect)
}
