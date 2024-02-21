package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, a []int, ops [][]int, expect []int) {
	res := solve(len(a), a, ops)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v , but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{1, 2, 3, 4}
	ops := [][]int{
		{2, 3},
		{1, 2, 2},
		{2, 1},
	}
	expect := []int{3, 2, 3, 4}
	runSample(t, a, ops, expect)
}

func TestSample2(t *testing.T) {
	a := []int{3, 50, 2, 1, 10}
	ops := [][]int{
		{1, 2, 0},
		{2, 8},
		{1, 3, 20},
	}
	expect := []int{8, 8, 20, 8, 10}
	runSample(t, a, ops, expect)
}
