package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, E [][]int, a []int, Q [][]int, expect []int) {
	res := solve(n, E, a, Q)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 7
	a := []int{1, 1, 1, 1, 1, 1, 1}
	E := [][]int{
		{1, 2},
		{1, 3},
		{2, 4},
		{2, 5},
		{3, 6},
		{6, 7},
	}
	Q := [][]int{
		{1, 6},
		{2, 3},
		{1, 6},
		{1, 2},
	}
	expect := []int{2, 3, 3}
	runSample(t, n, E, a, Q, expect)
}

func TestSample2(t *testing.T) {
	n := 10
	a := []int{-160016413, -90133231, -671446275, -314847579, -910548234, 121155052, -359359950, 83112406, -704889624, 145489303}
	E := [][]int{
		{1, 6},
		{1, 10},
		{10, 8},
		{1, 4},
		{3, 4},
		{2, 7},
		{2, 5},
		{3, 2},
		{9, 8},
	}
	Q := [][]int{
		{1, 4},
		{2, 2},
		{2, 4},
		{1, 4},
		{1, 10},
		{2, 10},
		{1, 9},
		{1, 6},
		{2, 8},
		{2, 10},
		{1, 5},
		{1, 8},
		{1, 1},
		{2, 5},
	}
	expect := []int{-2346335269, -314847579, -476287915, -704889624, 121155052, -1360041415, 228601709, -2861484545}
	runSample(t, n, E, a, Q, expect)
}
