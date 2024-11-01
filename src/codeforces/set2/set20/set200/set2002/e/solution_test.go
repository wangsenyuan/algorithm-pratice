package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, pairs [][]int, expect []int) {
	res := solve(pairs)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	pairs := [][]int{
		{2, 0},
		{1, 1},
		{3, 0},
		{5, 1},
	}
	expect := []int{2, 2, 4, 5}
	runSample(t, pairs, expect)
}

func TestSample2(t *testing.T) {
	pairs := [][]int{
		{4, 6},
		{1, 3},
		{4, 6},
		{4, 0},
		{7, 6},
		{6, 3},
	}
	expect := []int{4, 4, 7, 7, 10, 10}
	runSample(t, pairs, expect)
}

func TestSample3(t *testing.T) {
	pairs := [][]int{
		{9, 0},
		{7, 1},
		{5, 0},
		{7, 1},
		{9, 0},
		{1, 1},
		{2, 0},
	}
	expect := []int{9, 9, 9, 9, 9, 9, 10}
	runSample(t, pairs, expect)
}
