package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, E [][]int, A []int, expect []int) {
	res := solve(n, E, A)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 5
	E := [][]int{
		{2, 4},
		{3, 1},
		{3, 4},
		{3, 5},
	}
	A := []int{4, 5}
	expect := []int{2, 1, 2, 0, 0}
	runSample(t, n, E, A, expect)
}

func TestSample2(t *testing.T) {
	n := 8
	E := [][]int{
		{4, 1},
		{8, 4},
		{4, 5},
		{6, 4},
		{2, 5},
		{4, 3},
		{1, 7},
	}
	A := []int{2, 8, 3}
	expect := []int{3, 0, 0, 3, 1, 2, 3, 0}
	runSample(t, n, E, A, expect)
}

func TestSample3(t *testing.T) {
	n := 10
	E := [][]int{
		{2, 5},
		{4, 3},
		{7, 3},
		{7, 2},
		{5, 8},
		{3, 6},
		{8, 10},
		{7, 9},
		{7, 1},
	}
	A := []int{10, 6, 9, 1}
	expect := []int{0, 2, 2, 2, 2, 0, 2, 2, 0, 0}
	runSample(t, n, E, A, expect)
}
