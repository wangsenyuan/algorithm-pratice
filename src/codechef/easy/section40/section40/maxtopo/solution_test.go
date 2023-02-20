package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, E [][]int, k int, expect []int) {
	res := solve(n, E, k)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 4
	k := 2
	E := [][]int{
		{1, 2},
		{2, 3},
		{2, 4},
	}
	expect := []int{4, 2}
	runSample(t, n, E, k, expect)
}
