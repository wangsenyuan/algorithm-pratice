package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, qs [][]int, expect []int) {
	res := solve(qs)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	qs := [][]int{
		{4, 0},
		{3, 1},
		{5, 2},
		{6, 2},
		{18, 4},
		{10000, 3},
	}
	expect := []int{3, 1, 1, 2, 2, 22}
	runSample(t, qs, expect)
}
