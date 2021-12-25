package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, Q int, X []int, expect [][]int) {
	res := solve(Q, X)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %d %v, expect %v, but got %v", Q, X, expect, res)
	}
}

func TestSample(t *testing.T) {
	Q := 3
	X := []int{4, 2, 7}
	expect := [][]int{
		{0, 1},
		{1, 2},
		{3, 4},
	}
	runSample(t, Q, X, expect)
}
