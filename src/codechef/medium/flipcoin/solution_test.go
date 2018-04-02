package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, ops [][]int, expect []int) {
	res := solve(n, ops)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("sample %d %v, expect %v, but got %v", n, ops, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 4
	ops := [][]int{
		{1, 0, 3},
		{0, 1, 2},
		{1, 0, 1},
		{1, 0, 0},
		{0, 0, 3},
		{1, 0, 3},
		{1, 3, 3},
	}

	expect := []int{0, 1, 0, 2, 1}
	runSample(t, n, ops, expect)
}
