package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, C [][]int, Q []int, expect []int) {
	res := solve(len(C), C, Q)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	C := [][]int{
		{1, 3, 5},
		{2, 3, 5},
		{1, 2, 4},
		{2, 2, 5},
		{3, 2, 5},
	}
	Q := []int{1, 2, 3}
	expect := []int{3, 1, 1}
	runSample(t, C, Q, expect)
}
