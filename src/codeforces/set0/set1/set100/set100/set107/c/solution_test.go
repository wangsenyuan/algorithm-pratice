package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, y int64, P [][]int, expect []int) {
	res := solve(n, y, P)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	var y int64 = 2001
	P := [][]int{
		{1, 2},
		{1, 3},
	}
	expect := []int{1, 2, 3}
	runSample(t, n, y, P, expect)
}

func TestSample2(t *testing.T) {
	n := 7
	var y int64 = 2020
	P := [][]int{
		{1, 2},
		{1, 3},
		{2, 4},
		{2, 5},
		{3, 6},
		{3, 7},
	}
	expect := []int{1, 2, 3, 7, 4, 6, 5}
	runSample(t, n, y, P, expect)
}

func TestSample3(t *testing.T) {
	n := 10
	var y int64 = 3630801
	P := [][]int{}
	runSample(t, n, y, P, nil)
}

func TestSample4(t *testing.T) {
	n := 3
	var y int64 = 2001
	P := [][]int{
		{1, 2},
		{2, 3},
		{3, 1},
	}
	runSample(t, n, y, P, nil)
}
