package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, a []int, queries [][]int, b []int, expect []int) {
	res := solve(a, queries, b)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %v %v, expect %v, but got %v", a, queries, expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6}
	queries := [][]int{
		{2, 1, 3},
		{2, 3, 6},
		{1, 1, 6},
	}
	b := []int{2, 2, 1, 5, 3}
	expect := []int{3, 3, 1, 5, 2}

	runSample(t, a, queries, b, expect)
}

func TestSample2(t *testing.T) {
	a := []int{64, 3, 4, 665, 2}
	queries := [][]int{
		{1, 1, 3},
		{2, 1, 5},
	}
	b := []int{1, 2, 3, 4, 5}
	expect := []int{2, 665, 3, 64, 4}

	runSample(t, a, queries, b, expect)
}
