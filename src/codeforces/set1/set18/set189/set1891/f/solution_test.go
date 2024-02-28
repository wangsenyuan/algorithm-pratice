package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, ops [][]int, expect []int) {
	res := solve(ops)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	ops := [][]int{
		{2, 1, 3},
		{1, 1},
		{2, 2, 1},
		{1, 1},
		{2, 3, 2},
		{1, 3},
		{2, 1, 4},
		{1, 3},
		{2, 3, 2},
	}
	expect := []int{7, 5, 8, 6, 2}

	runSample(t, ops, expect)
}

func TestSample2(t *testing.T) {
	ops := [][]int{
		{2, 1, 1},
		{1, 1},
		{2, 1, -1},
		{1, 1},
		{2, 1, 1},
	}
	expect := []int{1, 0, 1}

	runSample(t, ops, expect)
}

func TestSample3(t *testing.T) {
	ops := [][]int{
		{1, 1},
		{1, 1},
		{2, 1, 1},
		{2, 1, 3},
		{2, 2, 10},
	}
	expect := []int{4, 14, 4}

	runSample(t, ops, expect)
}
