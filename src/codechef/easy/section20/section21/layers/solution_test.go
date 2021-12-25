package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n, c int, rects [][]int, expect []int64) {
	res := solve(n, c, rects)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, c := 5, 4
	rects := [][]int{
		{2, 2, 1},
		{2, 8, 3},
		{10, 2, 1},
		{8, 4, 2},
		{4, 6, 4},
	}
	expect := []int64{4, 16, 4, 24}
	runSample(t, n, c, rects, expect)
}

func TestSample2(t *testing.T) {
	n, c := 5, 4
	rects := [][]int{
		{2, 2, 1},
		{4, 8, 3},
		{10, 2, 1},
		{8, 4, 2},
		{4, 6, 4},
	}
	expect := []int64{4, 16, 8, 24}
	runSample(t, n, c, rects, expect)
}

func TestSample3(t *testing.T) {
	n, c := 5, 4
	rects := [][]int{
		{2, 2, 1},
		{6, 8, 3},
		{10, 2, 1},
		{8, 4, 2},
		{4, 6, 4},
	}
	expect := []int64{4, 16, 16, 24}
	runSample(t, n, c, rects, expect)
}

func TestSample4(t *testing.T) {
	n, c := 5, 4
	rects := [][]int{
		{2, 2, 1},
		{6, 8, 3},
		{10, 2, 1},
		{8, 4, 2},
		{4, 4, 4},
	}
	expect := []int64{4, 16, 24, 16}
	runSample(t, n, c, rects, expect)
}
