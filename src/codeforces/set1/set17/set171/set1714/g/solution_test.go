package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, E [][]int, expect []int) {
	res := solve(n, E)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 9
	E := [][]int{
		{1, 5, 6},
		{4, 5, 1},
		{2, 9, 10},
		{4, 2, 1},
		{1, 2, 1},
		{2, 3, 3},
		{6, 4, 3},
		{8, 1, 3},
	}
	expect := []int{0, 3, 1, 2, 1, 1, 2, 3}
	runSample(t, n, E, expect)
}

func TestSample2(t *testing.T) {
	n := 10
	E := [][]int{
		{1, 1, 4},
		{2, 3, 5},
		{2, 5, 1},
		{3, 4, 3},
		{3, 1, 5},
		{5, 3, 5},
		{5, 2, 1},
		{1, 3, 2},
		{6, 2, 1},
	}
	expect := []int{0, 1, 2, 1, 1, 2, 2, 1, 1}
	runSample(t, n, E, expect)
}
