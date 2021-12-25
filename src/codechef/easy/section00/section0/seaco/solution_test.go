package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, m int, cmds [][]int, expect []int64) {
	res := solve(n, m, cmds)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("sample %d %d %v, expects %v, but got %v", n, m, cmds, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m := 5, 5
	cmds := [][]int{
		{1, 1, 2},
		{1, 4, 5},
		{2, 1, 2},
		{2, 1, 3},
		{2, 3, 4},
	}

	runSample(t, n, m, cmds, []int64{7, 7, 0, 7, 7})
}

func TestSample2(t *testing.T) {
	n, m := 1, 2
	cmds := [][]int{
		{1, 1, 1},
		{1, 1, 1},
	}

	runSample(t, n, m, cmds, []int64{2})
}

func TestSample3(t *testing.T) {
	n, m := 10, 10
	cmds := [][]int{
		{1, 1, 10},
		{2, 1, 1},
		{2, 1, 2},
		{2, 1, 3},
		{2, 1, 4},
		{2, 1, 5},
		{2, 1, 6},
		{2, 1, 7},
		{2, 1, 8},
		{2, 1, 9},
	}
	runSample(t, n, m, cmds, []int64{512, 512, 512, 512, 512, 512, 512, 512, 512, 512})
}
