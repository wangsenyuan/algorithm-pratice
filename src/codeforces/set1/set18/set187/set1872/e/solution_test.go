package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, a []int, s string, ops [][]int, expect []int) {
	res := solve(a, s, ops)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	s := "01000"
	ops := [][]int{
		{2, 0},
		{2, 1},
		{1, 2, 4},
		{2, 0},
		{2, 1},
		{1, 1, 3},
		{2, 1},
	}
	expect := []int{3, 2, 6, 7, 7}
	runSample(t, a, s, ops, expect)
}
