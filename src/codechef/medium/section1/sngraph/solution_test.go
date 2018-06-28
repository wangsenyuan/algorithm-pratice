package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n, m int, edges [][]int, expect []int) {
	res := solve(n, m, edges)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %d %d %v, expect %v, but got %v", n, m, edges, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m := 5, 4
	edges := [][]int{
		{1, 2},
		{2, 3},
		{2, 4},
		{4, 5},
	}

	expect := []int{0, 3, 4, 4, 4}

	runSample(t, n, m, edges, expect)
}
