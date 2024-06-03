package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, cmds [][]int, expect []bool) {
	res := solve(cmds)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	cmds := [][]int{
		{1, 0},
		{1, 0},
		{1, 0},
		{2, 3},
		{2, 4},
	}
	expect := []bool{true, false}
	runSample(t, cmds, expect)
}

func TestSample2(t *testing.T) {
	cmds := [][]int{
		{1, 0},
		{1, 1},
		{1, 2},
		{1, 10},
		{2, 4},
		{2, 6},
		{2, 7},
	}
	expect := []bool{true, true, true}
	runSample(t, cmds, expect)
}
