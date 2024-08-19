package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, cmds [][]int, expect []int) {
	res := solve(cmds)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	cmds := [][]int{
		{1, 5},
		{0, 10},
		{1, -5},
		{0, 5},
		{1, 11},
		{0, -10},
	}
	expect := []int{5, 25, 10, 15, 36, 21}
	runSample(t, cmds, expect)
}
