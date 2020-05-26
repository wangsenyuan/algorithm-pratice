package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n, k int, G [][]int, expect [][]int) {
	res := solve(n, k, G)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %d %d %v, expect %v, but got %v", n, k, G, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, k := 4, 2
	G := [][]int{
		{0, 1, 2, 3},
		{4, 5, 6, 7},
		{8, 9, 0, 1},
		{2, 3, 4, 0},
	}
	expect := [][]int{
		{0, 1, 2},
		{4, 0, 0},
		{2, 0, 0},
	}
	runSample(t, n, k, G, expect)
}
