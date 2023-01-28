package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, E [][]int, expect []int) {
	res := solve(n, E)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 5
	E := [][]int{
		{1, 2},
		{3, 4},
		{3, 2},
		{4, 2},
		{2, 5},
	}
	expect := []int{1, 4}
	runSample(t, n, E, expect)
}
