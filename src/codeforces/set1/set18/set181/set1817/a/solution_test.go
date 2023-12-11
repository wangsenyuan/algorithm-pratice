package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, a []int, Q [][]int, expect []int) {
	res := solve(a, Q)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{1, 2, 4, 3, 3, 5, 6, 2, 1}
	Q := [][]int{
		{1, 3},
		{1, 4},
		{2, 5},
		{6, 6},
		{3, 7},
		{7, 8},
		{1, 8},
		{8, 8},
	}
	expect := []int{3, 4, 3, 1, 4, 2, 7, 1}
	runSample(t, a, Q, expect)
}
