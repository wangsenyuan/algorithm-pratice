package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, C []int, E [][]int, expect []int) {
	res := solve(n, C, E)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 8
	C := []int{0, 1, 0, 0, 0, 0, 1, 0}
	E := [][]int{
		{8, 6},
		{2, 5},
		{7, 8},
		{6, 5},
		{4, 5},
		{6, 1},
		{7, 3},
	}
	expect := []int{0, 1, 1, 1, 1, 0, 1, 1}
	runSample(t, n, C, E, expect)
}
