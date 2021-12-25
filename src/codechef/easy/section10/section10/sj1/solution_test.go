package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, edges [][]int, V []uint64, M []uint64, expect []uint64) {
	res := solve(n, edges, V, M)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %d %v %v %v, expect %v, but got %v", n, edges, V, M, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 5
	edges := [][]int{
		{1, 2}, {2, 5}, {1, 3}, {3, 4},
	}
	V := []uint64{2, 3, 4, 6, 7}
	M := []uint64{1, 2, 3, 2, 10}
	expect := []uint64{0, 9}
	runSample(t, n, edges, V, M, expect)
}
