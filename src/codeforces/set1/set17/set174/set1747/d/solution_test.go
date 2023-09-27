package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, v []int, q [][]int, expect []int) {
	res := solve(v, q)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	v := []int{3, 0, 3, 3, 1, 2, 3}
	q := [][]int{
		{3, 4},
		{4, 6},
		{3, 7},
		{5, 6},
		{1, 6},
		{2, 2},
	}
	expect := []int{-1, 1, 1, -1, 2, 0}
	runSample(t, v, q, expect)
}
