package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, E [][]int, A []int, expect []int64) {
	res := solve(n, E, A)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	A := []int{10, 11, 12}
	E := [][]int{
		{1, 2},
		{1, 3},
	}
	expect := []int64{0, 0}
	runSample(t, n, E, A, expect)
}

func TestSample2(t *testing.T) {
	n := 5
	A := []int{1, 1, 9, 4, 4}
	E := [][]int{
		{1, 2},
		{2, 3},
		{3, 4},
		{3, 5},
	}
	expect := []int64{1, 0, 4, 4}
	runSample(t, n, E, A, expect)
}

func TestSample3(t *testing.T) {
	n := 7
	A := []int{3, 8, 7, 8, 3, 7, 7}
	E := [][]int{
		{3, 6},
		{3, 7},
		{2, 5},
		{1, 6},
		{5, 6},
		{4, 1},
	}
	expect := []int64{7, 7, 8, 11, 11, 8}
	runSample(t, n, E, A, expect)
}
