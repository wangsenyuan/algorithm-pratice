package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, m int, k int, G [][]int, expect []int) {
	res := solve(n, m, k, G)
	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %d %d %d %v, expect %v, but got %v", n, m, k, G, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m, k := 5, 4, 2
	G := [][]int{
		{1, 2, 5},
		{1, 3, 9},
		{2, 4, 0},
		{2, 5, 1},
	}
	expect := []int{
		1, 4, 8, 4, 5,
	}
	runSample(t, n, m, k, G, expect)
}
