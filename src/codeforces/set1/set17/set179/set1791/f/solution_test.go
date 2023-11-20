package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, a []int, ops [][]int, expect []int) {
	res := solve(a, ops)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{1, 420, 69, 1434, 2023}
	ops := [][]int{
		{1, 2, 3},
		{2, 2},
		{2, 3},
		{2, 4},
		{1, 2, 5},
		{2, 1},
		{2, 3},
		{2, 5},
	}
	expect := []int{6, 15, 1434, 1, 6, 7}
	runSample(t, a, ops, expect)
}
