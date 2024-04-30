package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, a []int, cmds [][]int, expect []bool) {
	res := solve(a, cmds)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{2, 6, 3}
	cmds := [][]int{
		{1, 1, 2, 2},
		{1, 1, 3, 3},
		{2, 1, 9},
		{1, 1, 3, 2},
	}
	expect := []bool{true, true, false}

	runSample(t, a, cmds, expect)
}

func TestSample2(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	cmds := [][]int{
		{1, 1, 4, 2},
		{2, 3, 6},
		{1, 1, 4, 2},
		{1, 1, 5, 2},
		{2, 5, 10},
		{1, 1, 5, 2},
	}
	expect := []bool{false, true, false, true}

	runSample(t, a, cmds, expect)
}
