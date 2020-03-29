package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, E [][]int, expect []int64) {
	res := solve(n, E)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %d %v, expect %v,  but got %v", n, E, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 7
	E := [][]int{
		{1, 3, 2},
		{5, 1, -4},
		{2, 5, 3},
		{2, 6, 2},
		{3, 5, 1},
		{4, 6, -3},
		{7, 4, -3},
	}
	expect := []int64{INF, INF, INF, -6, INF, -6, -6}
	runSample(t, n, E, expect)
}
