package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, k int, a []int, cmds [][]int, expect []int) {
	res := solve(k, a, cmds)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	k := 2
	a := []int{7, 4, 9, 9}
	cmds := [][]int{
		{1, 2, 8},
		{3, 1},
		{4, 2, 4},
	}
	expect := []int{24}
	runSample(t, k, a, cmds, expect)
}

func TestSample2(t *testing.T) {
	k := 3
	a := []int{7, 0, 8, 8, 7, 1, 5, 2}
	cmds := [][]int{
		{4, 3, 7},
		{2, 1},
		{3, 2},
		{4, 1, 6},
		{2, 3},
		{1, 5, 16},
		{4, 8, 8},
		{3, 0},
	}
	expect := []int{29, 22, 1}
	runSample(t, k, a, cmds, expect)
}
