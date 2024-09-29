package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, mat [][]int, roads [][]int, expect []int) {
	res := solve(mat, roads)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	mat := [][]int{
		{0, 5},
		{5, 0},
	}
	roads := [][]int{
		{1, 2, 3},
	}
	expect := []int{3}
	runSample(t, mat, roads, expect)
}

func TestSample2(t *testing.T) {
	mat := [][]int{
		{0, 4, 5},
		{4, 0, 9},
		{5, 9, 0},
	}
	roads := [][]int{
		{2, 3, 8},
		{1, 2, 1},
	}
	expect := []int{17, 12}
	runSample(t, mat, roads, expect)
}
