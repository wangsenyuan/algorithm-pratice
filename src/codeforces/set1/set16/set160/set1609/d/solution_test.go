package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, cond [][]int, expect []int) {
	res := solve(n, cond)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 7
	cond := [][]int{
		{1, 2},
		{3, 4},
		{2, 4},
		{7, 6},
		{6, 5},
		{1, 7},
	}
	expect := []int{1, 1, 3, 3, 3, 6}
	runSample(t, n, cond, expect)
}

func TestSample2(t *testing.T) {
	n := 10
	cond := [][]int{
		{1, 2},
		{2, 3},
		{3, 4},
		{1, 4},
		{6, 7},
		{8, 9},
		{8, 10},
		{1, 4},
	}
	expect := []int{1, 2, 3, 4, 5, 5, 6, 8}
	runSample(t, n, cond, expect)
}
