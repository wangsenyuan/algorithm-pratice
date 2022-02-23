package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, src *Tree, qs []*Tree, expect []bool) {
	res := solve(src, qs)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	src := CreateTree(5, []int{1, 2, 1, 2, 1}, [][]int{
		{1, 2, 0},
		{1, 3, 1},
		{2, 4, 1},
		{3, 5, 0},
	})
	qs := []*Tree{
		CreateTree(2, []int{1, 2}, [][]int{{1, 2, 0}}),
		CreateTree(2, []int{1, 1}, [][]int{{1, 2, 0}}),
	}
	expect := []bool{false, true}
	runSample(t, src, qs, expect)
}
