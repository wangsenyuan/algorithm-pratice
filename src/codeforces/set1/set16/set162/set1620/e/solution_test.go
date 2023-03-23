package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, Q [][]int, expect []int) {
	res := solve(Q)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	Q := [][]int{
		{1, 3},
		{1, 1},
		{2, 1, 2},
		{1, 2},
		{1, 1},
		{1, 2},
		{2, 1, 3},
	}

	expect := []int{3, 2, 2, 3, 2}
	runSample(t, Q, expect)
}

func TestSample2(t *testing.T) {
	Q := [][]int{
		{1, 3},
		{1, 1},
		{2, 1, 2},
		{1, 1},
		{2, 2, 1},
		{1, 2},
		{1, 1},
		{1, 2},
		{2, 1, 3},
	}
	// 3, 1, 1, 2, 1, 2
	// 3, 3, 3, 2, 3, 2
	expect := []int{3, 3, 3, 2, 3, 2}
	runSample(t, Q, expect)
}

func TestSample3(t *testing.T) {
	Q := [][]int{
		{2, 1, 4},
		{1, 1},
		{1, 4},
		{1, 2},
		{2, 2, 4},
		{2, 4, 3},
		{1, 2},
		{2, 2, 7},
	}
	// 1, 3, 3, 7
	expect := []int{1, 3, 3, 7}
	runSample(t, Q, expect)
}
