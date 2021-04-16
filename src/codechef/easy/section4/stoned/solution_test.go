package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, H []int, Q [][]int, expect []int) {
	res := solve(H, Q)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	H := []int{3, 10, 20, 20, 5, 50}
	Q := [][]int{
		{1, 2},
		{2, 4},
		{4, 3},
		{4, 5},
		{1, 3},
		{4, 4},
	}
	expect := []int{2, 1, 1, 1, 1, 1}
	runSample(t, H, Q, expect)
}
