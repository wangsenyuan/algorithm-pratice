package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, roads [][]int, path []int, expect []int) {
	res := solve(n, roads, path)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 6
	roads := [][]int{
		{1, 5},
		{5, 4},
		{1, 2},
		{2, 3},
		{3, 4},
		{4, 1},
		{2, 6},
		{6, 4},
		{4, 2},
	}
	path := []int{1, 2, 3, 4}
	expect := []int{1, 2}

	runSample(t, n, roads, path, expect)
}

func TestSample2(t *testing.T) {
	n := 7
	roads := [][]int{
		{1, 2},
		{2, 3},
		{3, 4},
		{4, 5},
		{5, 6},
		{6, 7},
		{7, 1},
	}
	path := []int{1, 2, 3, 4, 5, 6, 7}
	expect := []int{0, 0}

	runSample(t, n, roads, path, expect)
}

func TestSample3(t *testing.T) {
	n := 8
	roads := [][]int{
		{8, 7},
		{8, 6},
		{7, 5},
		{7, 4},
		{6, 5},
		{6, 4},
		{5, 3},
		{5, 2},
		{4, 3},
		{4, 2},
		{3, 1},
		{2, 1},
		{1, 8},
	}
	path := []int{8, 7, 5, 2, 1}
	expect := []int{0, 3}

	runSample(t, n, roads, path, expect)
}
