package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, A []int, E [][]int, expect []int) {
	res := solve(n, A, E)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 5
	E := [][]int{
		{1, 2},
		{2, 3},
		{2, 4},
		{1, 5},
	}
	A := []int{2, 1, 3, 2, 1}
	expect := []int{0, 2, 1, 2}
	runSample(t, n, A, E, expect)
}

func TestSample2(t *testing.T) {
	n := 6
	E := [][]int{
		{1, 2},
		{1, 3},
		{1, 4},
		{4, 5},
		{4, 6},
	}
	A := []int{1, 2, 3, 1, 4, 5}
	expect := []int{1, 1, 0, 1, 1}
	runSample(t, n, A, E, expect)
}
