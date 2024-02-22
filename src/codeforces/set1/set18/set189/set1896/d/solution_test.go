package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, a []int, ops [][]int, expect []bool) {
	res := solve(a, ops)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{2, 1, 2, 1, 2}
	ops := [][]int{
		{1, 5},
		{1, 6},
		{1, 7},
		{2, 4, 2},
		{1, 7},
	}
	expect := []bool{true, true, false, true}
	runSample(t, a, ops, expect)
}

func TestSample2(t *testing.T) {
	a := []int{2, 2, 2}
	ops := [][]int{
		{1, 6},
		{1, 5},
	}
	expect := []bool{true, false}
	runSample(t, a, ops, expect)
}
