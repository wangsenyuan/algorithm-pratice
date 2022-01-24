package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, A []int, cmds [][]int, expect []int64) {
	res := solve(n, A, cmds)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	A := []int{1, 2, 2}
	cmds := [][]int{
		{2, 2},
		{2, 3},
	}
	expect := []int64{4, 8}
	runSample(t, n, A, cmds, expect)
}

func TestSample2(t *testing.T) {
	n := 5
	A := []int{1, 2, 3, 4, 5}
	cmds := [][]int{
		{2, 5},
		{1, 1, 2},
		{2, 3},
		{2, 4},
		{2, 5},
	}
	expect := []int64{35, 8, 17, 31}
	runSample(t, n, A, cmds, expect)
}
