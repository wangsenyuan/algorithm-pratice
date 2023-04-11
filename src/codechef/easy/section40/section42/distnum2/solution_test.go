package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, A []int, Q [][]int, expect []int) {
	res := solve(A, Q)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{3, 2, 1, 2}
	Q := [][]int{
		{0, 1, 0, 3, 2},
		{2, 0, 0, 3, 4},
		{1, 2, 1, 3, 2},
		{2, 0, 0, 3, 3},
	}
	expect := []int{2, -1, 2, 3}
	runSample(t, A, Q, expect)
}

func TestSample2(t *testing.T) {
	A := []int{9, 10, 6, 3, 8, 4, 9, 6, 4, 10}
	Q := [][]int{
		{0, 2, 0, 9, 3},
		{1, 9, 1, 3, 3},
		{1, 8, 1, 0, 3},
		{1, 2, 1, 7, 2},
		{1, 6, 1, 2, 3},
		{1, 4, 1, 3, 1},
		{1, 6, 1, 6, 1},
		{1, 4, 1, 8, 1},
		{1, 9, 1, 3, 3},
		{1, 9, 1, 2, 1},
	}
	expect := []int{6, 9, 10, 4, 6, 3, 10, 4, 6, 4}
	runSample(t, A, Q, expect)
}
