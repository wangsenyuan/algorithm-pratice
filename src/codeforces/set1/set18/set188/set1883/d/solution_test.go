package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, ops [][]int, expect []bool) {
	res := solve(ops)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	ops := [][]int{
		{0, 1, 2},
		{0, 3, 4},
		{0, 2, 3},
		{0, 2, 2},
		{0, 3, 4},
		{1, 3, 4},
		{1, 3, 4},
		{1, 1, 2},
		{0, 3, 4},
		{1, 2, 2},
		{1, 2, 3},
		{1, 3, 4},
	}
	expect := []bool{false, true, true, true, true, true, false, false, true, false, false, false}
	runSample(t, ops, expect)
}
