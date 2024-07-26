package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, buildings [][]int32, expect [][]int32) {
	res := solve(buildings)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	buildings := [][]int32{
		{3, 0, 2},
		{4, 1, 3},
	}
	expect := [][]int32{
		{0, 0},
		{0, 3},
		{1, 3},
		{1, 4},
		{3, 4},
		{3, 0},
	}
	runSample(t, buildings, expect)
}

func TestSample2(t *testing.T) {
	buildings := [][]int32{
		{3, -3, 0},
		{2, -1, 1},
		{4, 2, 4},
		{2, 3, 7},
		{3, 6, 8},
	}
	expect := [][]int32{
		{-3, 0},
		{-3, 3},
		{0, 3},
		{0, 2},
		{1, 2},
		{1, 0},
		{2, 0},
		{2, 4},
		{4, 4},
		{4, 2},
		{6, 2},
		{6, 3},
		{8, 3},
		{8, 0},
	}
	runSample(t, buildings, expect)
}

func TestSample3(t *testing.T) {
	buildings := [][]int32{
		{5, -5, -4},
		{3, -3, 0},
		{2, -1, 1},
		{1, 0, 1},
		{4, 2, 4},
		{2, 3, 7},
		{3, 6, 8},
	}
	expect := [][]int32{
		{-5, 0},
		{-5, 5},
		{-4, 5},
		{-4, 0},
		{-3, 0},
		{-3, 3},
		{0, 3},
		{0, 2},
		{1, 2},
		{1, 0},
		{2, 0},
		{2, 4},
		{4, 4},
		{4, 2},
		{6, 2},
		{6, 3},
		{8, 3},
		{8, 0},
	}
	runSample(t, buildings, expect)
}
