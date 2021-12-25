package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n, q int, E [][]int, A, V []int, expect []int64) {
	res := solve(n, q, A, E, V)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %d %d %v %v %v, expect %v, but got %v", n, q, E, A, V, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, q := 4, 3
	E := [][]int{
		{1, 2}, {2, 3}, {3, 4},
	}
	A := []int{6, 2, 7, 3}
	V := []int{3, 2, 1}
	expect := []int64{13, 5, 0, 0}
	runSample(t, n, q, E, A, V, expect)
}
