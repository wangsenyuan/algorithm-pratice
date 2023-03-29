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
	n := 4
	E := [][]int{
		{1, 2},
		{2, 3},
		{1, 4},
	}
	expect := []int{1, 2, 2, 2}
	runSample(t, n, E, expect)
}

func TestSample2(t *testing.T) {
	n := 2
	E := [][]int{
		{1, 2},
	}
	expect := []int{1, 1}
	runSample(t, n, E, expect)
}

func TestSample3(t *testing.T) {
	n := 7
	E := [][]int{
		{1, 2},
		{1, 3},
		{3, 4},
		{3, 5},
		{3, 6},
		{5, 7},
	}
	expect := []int{2, 2, 3, 3, 3, 3, 3}
	runSample(t, n, E, expect)
}
