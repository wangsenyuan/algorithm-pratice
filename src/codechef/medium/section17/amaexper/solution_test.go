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
		{1, 2, 5},
		{1, 5, 6},
		{2, 3, 2},
		{2, 4, 1},
		{3, 6, 1},
	}
	expect := []int{8, 3, 1, 0, 0, 0}
	runSample(t, n, E, expect)
}
