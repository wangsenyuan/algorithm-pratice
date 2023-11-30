package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, s string, Q [][]int, expect []bool) {
	res := solve(s, Q)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "tedubcyyxopz"
	Q := [][]int{
		{2, 2, 5},
		{1, 4, 8, 2},
		{2, 1, 7},
		{1, 3, 5, 40},
		{1, 9, 11, 3},
		{1, 10, 10, 9},
		{2, 4, 10},
		{2, 10, 12},
	}
	expect := []bool{true, false, false, true}
	runSample(t, s, Q, expect)
}

func TestSample2(t *testing.T) {
	s := "ubnxwwgzjt"
	Q := [][]int{
		{2, 4, 10},
		{2, 10, 10},
		{1, 6, 10, 8},
		{2, 7, 7},
	}
	expect := []bool{false, true, true}
	runSample(t, s, Q, expect)
}

func TestSample3(t *testing.T) {
	s := "hntcxfxyhtu"
	Q := [][]int{
		{1, 4, 6, 1},
		{2, 4, 10},
		{1, 4, 10, 21},
	}
	expect := []bool{true}
	runSample(t, s, Q, expect)
}
