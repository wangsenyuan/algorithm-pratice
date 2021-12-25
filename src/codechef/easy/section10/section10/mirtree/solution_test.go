package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n, q int, ops [][]int, expect []int) {
	res := solve(n, q, ops)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("sample %d %d %v, expect %v, but got %v", n, q, ops, expect, res)
	}
}

func TestSample1(t *testing.T)  {
	n, q := 2, 3
	ops := [][]int{
		{1, 1},
		{1, 4},
		{2},
	}
	expect := []int{38}
	runSample(t, n, q, ops, expect)
}