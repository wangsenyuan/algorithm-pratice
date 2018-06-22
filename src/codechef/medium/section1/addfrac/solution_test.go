package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, A [][]int, expect [][]int64) {
	res := solve(n, A)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %d %v, expect %v, but got %v", n, A, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 4
	A := [][]int{
		{1, 1},
		{4, 3},
		{10, 1},
		{5, 4},
	}
	expect := [][]int64{
		{3, 1},
		{7, 2},
		{10, 1},
		{5, 4},
	}
	runSample(t, n, A, expect)
}

func TestSample2(t *testing.T) {
	n := 5
	A := [][]int{
		{1, 3},
		{3, 1},
		{2, 5},
		{5, 6},
		{6, 5},
	}
	expect := [][]int64{
		{1, 1},
		{3, 1},
		{13, 16},
		{1, 1},
		{6, 5},
	}

	runSample(t, n, A, expect)
}
