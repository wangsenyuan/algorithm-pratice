package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, k int, P [][]int, C []int, expect [][]int) {
	res := solve(n, k, P, C)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, k := 4, 11
	C := []int{4, 2, 3, 1}
	P := [][]int{
		{0, 1},
		{0, 2},
		{1, 2},
		{2, 3},
		{3, 1},
	}
	expect := [][]int{
		{0, 1, 2},
		{0, 2},
		{0, 1},
		{0},
		{1, 2, 3},
		{1, 2},
		{2, 3},
		{2},
		{1, 3},
		{1},
		{3},
	}
	runSample(t, n, k, P, C, expect)
}


func TestSample2(t *testing.T) {
	n, k := 1, 1
	C := []int{1000000000}

	expect := [][]int{
		{0},
	}
	runSample(t, n, k, nil, C, expect)
}
