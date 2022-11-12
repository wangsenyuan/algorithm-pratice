package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, E [][]int, Q [][]int, expect []int) {
	res := solve(n, E, Q)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	E := [][]int{
		{1, 2},
		{2, 3},
	}
	Q := [][]int{{1, 2}}
	expect := []int{3}
	runSample(t, n, E, Q, expect)
}

func TestSample2(t *testing.T) {
	n := 5
	E := [][]int{
		{1, 2},
		{5, 4},
		{1, 5},
		{5, 3},
	}
	Q := [][]int{{1, 2}, {6, 1}}
	expect := []int{5, 4}
	runSample(t, n, E, Q, expect)
}
