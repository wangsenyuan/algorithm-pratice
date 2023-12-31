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
	a := []int{8, 1, 6, 3, 7}
	Q := [][]int{
		{2, 1, 5},
		{2, 3, 4},
		{1, 1, 3},
		{2, 3, 4},
	}
	expect := []int{10, 2, 1}
	runSample(t, a, Q, expect)
}
