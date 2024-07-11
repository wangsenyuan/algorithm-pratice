package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, m int, k int, x [][]int, y [][]int, expect [][]int) {
	res := solve(n, m, k, x, y)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m := 3, 3
	k := 10
	x := [][]int{
		{1, 1},
		{1, 1},
		{1, 1},
	}
	y := [][]int{
		{1, 1, 1},
		{1, 1, 1},
	}
	expect := [][]int{
		{10, 10, 10},
		{10, 10, 10},
		{10, 10, 10},
	}

	runSample(t, n, m, k, x, y, expect)
}
