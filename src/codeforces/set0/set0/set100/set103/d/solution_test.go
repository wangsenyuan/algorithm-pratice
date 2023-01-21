package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, W []int, Q [][]int, expect []int64) {
	res := solve(W, Q)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	W := []int{1, 2, 3}
	Q := [][]int{
		{1, 1},
		{1, 2},
	}
	expect := []int64{6, 4}
	runSample(t, W, Q, expect)
}

func TestSample2(t *testing.T) {
	W := []int{2, 3, 5, 7}
	Q := [][]int{
		{1, 3},
		{2, 3},
		{2, 2},
	}
	expect := []int64{9, 3, 10}
	runSample(t, W, Q, expect)
}
