package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, E [][]int, expect []int) {
	res := solve(n, E)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 6
	E := [][]int{
		{1, 2},
		{1, 3},
		{2, 4},
		{2, 5},
		{3, 6},
	}
	expect := []int{1, 1, 2, 4, 6, 6}
	runSample(t, n, E, expect)
}

func TestSample2(t *testing.T) {
	n := 5
	E := [][]int{
		{1, 2},
		{2, 3},
		{3, 4},
		{3, 5},
	}
	expect := []int{1, 1, 3, 5, 5}
	runSample(t, n, E, expect)
}
