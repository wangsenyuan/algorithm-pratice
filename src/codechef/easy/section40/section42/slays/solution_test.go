package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, S []int, Q [][]int, expect []int) {
	res := solve(S, Q)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	S := []int{1, 3, 2, 2, 1}
	Q := [][]int{
		{2, 4},
		{3, 4},
		{4, 5},
	}
	expect := []int{4, 3, 2}
	runSample(t, S, Q, expect)
}

func TestSample2(t *testing.T) {
	S := []int{5, 6, 3, 5, 4, 4, 3, 2, 1}
	Q := [][]int{
		{1, 3},
		{2, 6},
		{3, 7},
		{6, 9},
	}
	expect := []int{6, 7, 6, 4}
	runSample(t, S, Q, expect)
}
