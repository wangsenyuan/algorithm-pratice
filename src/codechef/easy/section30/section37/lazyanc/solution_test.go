package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, E [][]int, A []int, expect []int64) {
	res := solve(n, E, A)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 5
	E := [][]int{
		{1, 2},
		{2, 3},
		{1, 4},
		{2, 5},
	}
	A := []int{3, 6, 8, 2, 1}
	expect := []int64{9, 11, 11, 5, 6}
	runSample(t, n, E, A, expect)
}
