package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n, m int, E [][]int, expect []int64) {
	res := solve(n, m, E)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m := 5, 6
	E := [][]int{
		{1, 2, 3},
		{2, 3, 4},
		{3, 4, 5},
		{4, 5, 6},
		{1, 5, 1},
		{2, 4, 2},
	}
	expect := []int64{0, 98, 49, 25, 114}

	runSample(t, n, m, E, expect)
}

func TestSample2(t *testing.T) {
	n, m := 3, 2
	E := [][]int{
		{1, 2, 1},
		{2, 3, 2},
	}
	expect := []int64{0, -1, 9}

	runSample(t, n, m, E, expect)
}

func TestSample3(t *testing.T) {
	n, m := 10, 20
	E := [][]int{
		{10, 1, 15},
		{7, 1, 32},
		{5, 3, 36},
		{3, 9, 14},
		{3, 4, 19},
		{6, 8, 4},
		{9, 6, 18},
		{7, 3, 38},
		{10, 7, 12},
		{7, 5, 29},
		{7, 6, 14},
		{6, 2, 40},
		{8, 9, 19},
		{7, 8, 11},
		{7, 4, 19},
		{2, 1, 38},
		{10, 9, 3},
		{6, 5, 50},
		{10, 3, 41},
		{1, 8, 3},
	}
	expect := []int64{0, 2201, 779, 1138, 1898, 49, 196, 520, 324, 490}

	runSample(t, n, m, E, expect)
}
