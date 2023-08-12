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
		{5, 3, 4},
		{2, 1, 1},
		{3, 2, 2},
		{2, 4, 2},
	}
	expect := []int64{1, 2, 2, 4}
	runSample(t, n, E, expect)
}

func TestSample2(t *testing.T) {
	n := 6
	E := [][]int{
		{3, 1, 1},
		{3, 6, 2},
		{5, 4, 2},
		{4, 2, 2},
		{6, 1, 1},
		{5, 2, 1},
		{3, 2, 3},
		{1, 5, 4},
	}
	expect := []int64{2, 1, 4, 3, 1}
	runSample(t, n, E, expect)
}
