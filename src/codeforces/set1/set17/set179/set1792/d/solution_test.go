package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, a [][]int, expect []int) {
	res := solve(len(a), a)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := [][]int{
		{2, 4, 1, 3},
		{1, 2, 4, 3},
		{2, 1, 3, 4},
	}
	expect := []int{1, 4, 4}
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := [][]int{
		{1, 2},
		{2, 1},
	}
	expect := []int{2, 2}
	runSample(t, a, expect)
}

func TestSample3(t *testing.T) {
	a := [][]int{
		{3, 4, 9, 6, 10, 2, 7, 8, 1, 5},
		{3, 9, 1, 8, 5, 7, 4, 10, 2, 6},
		{3, 10, 1, 7, 5, 9, 6, 4, 2, 8},
		{1, 2, 3, 4, 8, 6, 10, 7, 9, 5},
		{1, 2, 3, 4, 10, 6, 8, 5, 7, 9},
		{9, 6, 1, 2, 10, 4, 7, 8, 3, 5},
		{7, 9, 3, 2, 5, 6, 4, 8, 1, 10},
		{9, 4, 3, 7, 5, 6, 1, 10, 8, 2},
	}
	expect := []int{10, 8, 1, 6, 8, 10, 1, 7}
	runSample(t, a, expect)
}
