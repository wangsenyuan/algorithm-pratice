package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, A []int64, E [][]int, expect []int64) {
	res := solve(n, A, E)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	A := []int64{4, 5, 6}
	E := [][]int{
		{1, 2},
		{1, 3},
	}
	expect := []int64{1, -1, -1}
	runSample(t, n, A, E, expect)
}
